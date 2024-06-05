package data

import (
	"context"
	"database/sql"
	"go-starter/internal/config"
	"go-starter/internal/data/ent"
	"go-starter/internal/data/ent/migrate"
	"go-starter/pkg/log"

	"github.com/meilisearch/meilisearch-go"
	"github.com/redis/go-redis/v9"

	entsql "entgo.io/ent/dialect/sql"

	// mysql
	_ "github.com/go-sql-driver/mysql"
	// postgres
	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	err error
	db  *sql.DB
	ec  *ent.Client
	rc  *redis.Client
	ms  *meilisearch.Client
)

// Data .
type Data struct {
	db *sql.DB
	ec *ent.Client
	rc *redis.Client
	ms *meilisearch.Client
}

// New creates a new Database Connection.
func New(conf *config.Data) (*Data, func(), error) {
	ec, db := newClient(&conf.Database)
	d := &Data{
		db: db,
		ec: ec,
		rc: newRedis(&conf.Redis),
		// ms: newMeilisearch(&conf.Meilisearch),
	}

	cleanup := func() {
		log.Printf(context.Background(), "execute data cleanup.")
		if errs := d.Close(); errs != nil {
			log.Errorf(context.Background(), "cleanup errors: %v", errs)
		}
	}

	return d, cleanup, nil
}

// newRedis creates a new Redis datastore.
func newRedis(conf *config.Redis) *redis.Client {
	if conf == nil {
		log.Fatalf(context.Background(), "redis configuration cannot be nil")
	}
	rc = redis.NewClient(&redis.Options{
		Addr:         conf.Addr,
		Username:     conf.Username,
		Password:     conf.Password,
		DB:           int(conf.Db),
		ReadTimeout:  conf.ReadTimeout,
		WriteTimeout: conf.WriteTimeout,
		DialTimeout:  conf.DialTimeout,
		PoolSize:     10,
	})

	timeout, cancelFunc := context.WithTimeout(context.Background(), conf.DialTimeout)
	defer cancelFunc()
	if err := rc.Ping(timeout).Err(); err != nil {
		log.Fatalf(context.Background(), "redis connect error: %s", err)
	}

	return rc
}

// GetRedis returns the redis client
func (d *Data) GetRedis() *redis.Client {
	return d.rc
}

// newClient creates a new ent client.
func newClient(conf *config.Database) (*ent.Client, *sql.DB) {
	// Open database
	switch conf.Driver {
	case "postgres":
		db, err = sql.Open("pgx", conf.Source)
	case "mysql":
		db, err = sql.Open("mysql", conf.Source)
	default:
		log.Fatalf(context.Background(), "Dialect %v not supported.", conf.Driver)
	}

	if err != nil {
		log.Fatalf(context.Background(), "Failed to connect to database %v", err)
		return nil, nil
	}

	// Idle Connection
	db.SetMaxIdleConns(conf.MaxIdleConn)
	// Max Open Connection
	db.SetMaxOpenConns(conf.MaxOpenConn)
	// Max Connect Lifetime
	db.SetConnMaxLifetime(conf.ConnMaxLifeTime)

	ec = ent.NewClient(ent.Driver(entsql.OpenDB(conf.Driver, db)))
	// Auto migrate
	if conf.Migrate {
		if err = ec.Schema.Create(context.Background(), migrate.WithForeignKeys(false), migrate.WithDropIndex(true), migrate.WithDropColumn(true)); err != nil {
			log.Fatalf(context.Background(), "data.client.Schema.Create error: %v", err)
		}
	}

	return ec, db
}

// GetEntClient returns the ent client
func (d *Data) GetEntClient() *ent.Client {
	return d.ec
}

// GetDB returns the database
func (d *Data) GetDB() *sql.DB {
	return d.db
}

// newMeilisearch creates a new Meilisearch client.
func newMeilisearch(conf *config.Meilisearch) *meilisearch.Client {
	if conf == nil || conf.Host == "" {
		log.Fatalf(nil, "Meilisearch configuration cannot be nil")
	}

	ms = meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   conf.Host,
		APIKey: conf.APIKey,
	})

	// Check connection
	_, err := ms.Health()
	if err != nil {
		log.Fatalf(nil, "Meilisearch connect error: %v", err)
	}

	return ms
}

// GetMeilisearch returns the Meilisearch client
func (d *Data) GetMeilisearch() *meilisearch.Client {
	return d.ms
}

// Close .
func (d *Data) Close() (errs []error) {
	if err := d.rc.Close(); err != nil {
		errs = append(errs, err)
	}
	if err := d.db.Close(); err != nil {
		errs = append(errs, err)
	}
	if err := d.ec.Close(); err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

// Ping .
func (d *Data) Ping(ctx context.Context) error {
	return d.db.PingContext(ctx)
}
