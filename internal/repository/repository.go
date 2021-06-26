package repository

import (
	"context"
	"database/sql"
	"haiyon/go-starter/internal/generated/ent"
	"haiyon/go-starter/pkg/conf"
	"haiyon/go-starter/pkg/database/orm"
)

// Repository .
type Repository struct {
	client *ent.Client
	db     *sql.DB
}

// New new a instance
func New(cfg *conf.Config) (r *Repository) {
	c, d := orm.New(&cfg.ORM)
	r = &Repository{
		client: c,
		db:     d,
	}
	r.initORM()
	return
}

func (r *Repository) initORM() {}

// Close close the resource.
func (r *Repository) Close() error {
	return r.client.Close()
}

// Ping ping
func (r *Repository) Ping(ctx context.Context) (err error) {
	err = r.db.PingContext(ctx)
	return
}
