package main

import (
	"context"
	"errors"
	"fmt"
	"go-starter/internal/config"
	"go-starter/internal/server"
	"go-starter/pkg/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "go-starter/docs"
)

// Version Version number, can be specified during compilation: go build -ldflags "-X main.Version=x.x.x"
var Version = "dev+"

// @title go-starter
// @version 0.1.0
// @description a modern content management system
// @termsOfService https://go-starter.com

func main() {
	log.SetVersion(Version)

	// Loading config
	if err := config.Init(); err != nil {
		log.Fatalf(context.Background(), "❌ Config initialization error: %+v", err)
	}

	// Initialize logger
	loggerClean, err := log.Init(config.G.Logger)
	if err != nil {
		log.Fatalf(context.Background(), "❌ Logger initialization error: %+v", err)
	}
	defer loggerClean()

	// Print application name
	log.Infof(context.Background(), "%s", config.G.AppName)

	// Create server
	handler, cleanup, err := server.New(config.G)
	if err != nil {
		log.Fatalf(context.Background(), "❌ Failed to start server: %+v", err)
	}

	// Cleanup
	defer cleanup()

	// Start HTTP server
	addr := fmt.Sprintf("%s:%d", config.G.Host, config.G.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	log.Infof(context.Background(), "🚀 Listening and serving HTTP on: %s", addr)

	go func() {
		// Service connections
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf(context.Background(), "listen: %s", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Infof(context.Background(), "⌛️ Shutting down server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf(context.Background(), "❌ Server shutdown:", err)
	}
	// Catching ctx.Done(). Timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Infof(context.Background(), "⌛️ Timeout of 3 seconds.")
	}
	log.Infof(context.Background(), "👋 Server exiting")
}
