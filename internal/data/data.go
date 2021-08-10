package data

import (
	"context"
	"database/sql"
	"go-starter/common/conf"
	"go-starter/common/log"
	"go-starter/internal/generated/ent"
	"go-starter/internal/generated/ent/migrate"

	entsql "entgo.io/ent/dialect/sql"
	// mysql
	_ "github.com/go-sql-driver/mysql"
	// postgres
	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	err    error
	client *ent.Client
	db     *sql.DB
)

type Data struct {
	Client *ent.Client
	DB     *sql.DB
}

// New New Database Connection
func New(db *conf.DBConfig) (*Data, error) {
	dt, err := createWithConnection(&conf.DBConfig{
		Dialect:         db.Dialect,
		AutoMigrate:     db.AutoMigrate,
		DSN:             db.DSN,
		TablePrefix:     db.TablePrefix,
		MaxIdleConn:     db.MaxIdleConn,
		MaxOpenConn:     db.MaxOpenConn,
		ConnMaxLifeTime: db.ConnMaxLifeTime,
	})
	if err != nil {
		log.Fatalf(context.Background(), "Failed create data: %v", err)
		return nil, err
	}

	return dt, nil
}

// createWithConnection Create Connection Database
func createWithConnection(cfg *conf.DBConfig) (*Data, error) {
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
		log.Fatalf(context.Background(), "Failed to connect to database %v", err)
		return nil, err
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
			log.Fatalf(context.Background(), "data.client.Schema.Create error: %v", err)
		}
	}

	return &Data{
		Client: client,
		DB:     db,
	}, nil
}
