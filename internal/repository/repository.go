package repository

import (
	"context"
	"database/sql"
	"go-starter/common/conf"
	"go-starter/internal/data"
	"go-starter/internal/generated/ent"
)

// Repository .
type Repository struct {
	Client *ent.Client
	DB     *sql.DB
}

// New new a instance
func New(cfg *conf.Config) (repo *Repository) {
	nd, _ := data.New(&cfg.DB)
	return &Repository{
		Client: nd.Client,
		DB:     nd.DB,
	}
}

// Close close the resource.
func (repo *Repository) Close() error {
	return repo.Client.Close()
}

// Ping ping
func (repo *Repository) Ping(ctx context.Context) (err error) {
	err = repo.DB.PingContext(ctx)
	return
}
