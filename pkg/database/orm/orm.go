package orm

import (
	"context"
	"database/sql"
	"haiyon/go-starter/internal/generated/ent"
	"haiyon/go-starter/internal/generated/ent/migrate"
	"haiyon/go-starter/pkg/conf"
	"haiyon/go-starter/pkg/log"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	// mysql
	_ "github.com/go-sql-driver/mysql"
	// postgres
	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	err error
)

// Config Connection Pool
type Config struct {
	Dialect         string        // db type
	AutoMigrate     bool          // auto migrate
	DSN             string        // data source name
	TablePrefix     string        // table prefix
	MaxIdleConn     int           // pool
	MaxOpenConn     int           // pool
	ConnMaxLifeTime time.Duration // connect max life time
}

// New New Database Connection
func New(orm *conf.ORMConfig) (*ent.Client, *sql.DB) {
	return createWithConnection(&Config{
		Dialect:         orm.Dialect,
		AutoMigrate:     orm.AutoMigrate,
		DSN:             orm.DSN,
		TablePrefix:     orm.TablePrefix,
		MaxIdleConn:     orm.MaxIdleConn,
		MaxOpenConn:     orm.MaxOpenConn,
		ConnMaxLifeTime: orm.ConnMaxLifeTime,
	})
}

// createWithConnection Create Connection Database
func createWithConnection(cfg *Config) (client *ent.Client, db *sql.DB) {
	// open database
	switch cfg.Dialect {
	case "postgres":
		db, err = sql.Open("pgx", cfg.DSN)
	case "mysql":
		db, err = sql.Open("mysql", cfg.DSN)
	default:
		log.Fatalf(context.Background(), "Dialect %v not supported.", cfg.Dialect)
	}

	if err != nil {
		log.Fatalf(context.Background(), "Failed to connect to database \n%v", err)
	}

	// Idle Connection
	db.SetMaxIdleConns(cfg.MaxIdleConn)
	// Max Open Connection
	db.SetMaxOpenConns(cfg.MaxOpenConn)
	// Max Connect Lifetime
	db.SetConnMaxLifetime(cfg.ConnMaxLifeTime)

	client = ent.NewClient(ent.Driver(entsql.OpenDB(cfg.Dialect, db)))

	// auto migrate
	if cfg.AutoMigrate {
		if err = client.Schema.Create(context.Background(), migrate.WithDropIndex(true), migrate.WithDropColumn(true)); err != nil {
			log.Fatalf(context.Background(), "orm.client.Schema.Create error: %v\n", err)
		}
	}

	return
}
