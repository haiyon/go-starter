package data

import (
	"context"
	"database/sql"
	"fmt"
	"go-starter/internal/config"
	"go-starter/internal/data/ent"
	"go-starter/internal/data/ent/migrate"
	"go-starter/pkg/log"

	"github.com/redis/go-redis/v9"

	entsql "entgo.io/ent/dialect/sql"

	// mysql
	_ "github.com/go-sql-driver/mysql"
	// postgres
	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	err error
	db  *sql.DB
	ec  *ent.Client
	rc  redis.Cmdable
)

// Data .
type Data struct {
	ec *ent.Client
	rc redis.Cmdable
	db *sql.DB
}

// New creates a new Database Connection.
func New(conf *config.Data) (*Data, func(), error) {
	ec, db := newClient(&conf.Database)
	d := &Data{
		ec: ec,
		rc: newRedis(&conf.Redis),
		db: db,
	}

	cleanup := func() {
		log.Printf(context.Background(), "execute data cleanup of content service.")
		if err := d.ec.Close(); err != nil {
			log.Errorf(context.Background(), err.Error())
		}
	}

	return d, cleanup, nil
}

type Cache[T any] interface {
	// getCache data from cache using the specified field and return a pointer to type T and a possible error.
	getCache(context.Context, string) (*T, error)
	// setCache Save the specified data into cache using the specified string as key.
	setCache(context.Context, *T, string)
	// deleteCache data from cache using the specified key and an optional filter condition.
	deleteCache(context.Context, string)
	// resetCache data in cache using the specified pointer to type T as new value.
	resetCache(context.Context, *T, string)
}

// newRedis creates a new Redis datastore.
func newRedis(conf *config.Redis) redis.Cmdable {
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
		log.Fatalf(context.Background(), "redis connect error: %w", err)
	}

	return rc
}

// getCacheKey - define cache key of the user service.
func getCacheKey(key string) string {
	return fmt.Sprintf("%s_%s:%s", "sc", "sample", key) // os.Args[0]
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
