/*
Package main
Swagger 文档规则请参考：https://github.com/swaggo/swag#declarative-comments-format
使用方式：
	go get -u github.com/swaggo/swag/cmd/swag
	swag init --generalInfo ./main.go --output ./swagger
	make swag
*/
package main

import (
	"context"
	"fmt"
	"go-starter/common/log"
	"go-starter/internal/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-starter/common/conf"

	_ "go-starter/swagger"
)

// Version 版本号，可以通过编译的方式指定版本号：go build -ldflags "-X main.Version=x.x.x"
var Version = "dev+"

// @title Go Starter
// @version 0.1.0
// @description Go Starter Project Layout
// @termsOfService https://domain.com

func main() {
	log.SetVersion(Version)

	// loading config
	if err := conf.Init(); err != nil {
		log.Fatalf(context.Background(), "❌ conf init error: %+v", err)
	}
	// init logger
	loggerClean, err := log.Init(conf.G.Logger)
	if err != nil {
		log.Fatalf(context.Background(), "❌ logger init error: %+v", err)
	}
	defer loggerClean()

	// print application name
	log.Infof(context.Background(), "%s", conf.G.AppName)

	// new a serve
	serve, svc := server.New(conf.G)
	if err != nil {
		log.Fatalf(context.Background(), "❌ Failed to run server: %+v", err)
	}

	// Start http server
	addr := fmt.Sprintf("%s:%d", conf.G.Host, conf.G.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: serve,
	}
	log.Infof(context.Background(), "🚀 Listening and serving HTTP on: %s", addr)

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf(context.Background(), "listen: %s", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// // kill (no param) default send syscanll.SIGTERM
	// // kill -2 is syscall.SIGINT
	// // kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Infof(context.Background(), "⌛️ Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf(context.Background(), "❌ Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		// log.Infof(context.Background(), "⌛️ timeout of 3 seconds.")
		log.Printf(context.Background(), "💡 Database Close...")
		if err := svc.DBClose(); err != nil {
			log.Printf(context.Background(), "❌ Database Close error:", err)
		}

	}
	log.Infof(context.Background(), "👋 Server exiting")
}
