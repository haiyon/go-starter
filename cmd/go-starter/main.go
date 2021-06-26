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
	"haiyon/go-starter/internal/usage/http"
	"haiyon/go-starter/pkg/log"

	"haiyon/go-starter/pkg/conf"

	_ "haiyon/go-starter/swagger"
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
		log.Fatalf(context.Background(), "conf init error: %+v", err)
	}
	// init logger
	loggerClean, err := log.Init(conf.G.Logger)
	if err != nil {
		log.Fatalf(context.Background(), "logger init error: %+v", err)
	}
	defer loggerClean()

	// print application name
	log.Infof(context.Background(), "%s\n", conf.G.AppName)

	// new a http
	if err := http.New(conf.G); err != nil {
		log.Fatalf(context.Background(), "Failed to run server: %+v", err)
	}
}
