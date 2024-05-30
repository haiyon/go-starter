package server

import (
	"context"
	"go-starter/internal/config"
	"go-starter/internal/data"
	"go-starter/internal/handler"
	"go-starter/internal/service"
	"go-starter/pkg/log"
	"net/http"
)

// New creates a new server.
func New(conf *config.Config) (http.Handler, func(), error) {
	d, cleanup, err := data.New(&conf.Data)
	if err != nil {
		log.Fatalf(context.Background(), "❌ Failed initializing data: %+v", err)
		// panic(err)
	}

	// Initialize services
	svc := service.New(conf, d)

	// New HTTP server
	h, err := newHTTP(conf, handler.New(svc), svc)
	if err != nil {
		log.Fatalf(context.Background(), "❌ Failed initializing http: %+v", err)
		// panic(err)
	}

	return h, cleanup, nil
}
