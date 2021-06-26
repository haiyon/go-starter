package http

import (
	"context"
	"fmt"
	"net/http"
	"haiyon/go-starter/internal/controllers"
	"haiyon/go-starter/internal/services"
	"haiyon/go-starter/internal/usage/middlewares"
	"haiyon/go-starter/pkg/conf"
	"haiyon/go-starter/pkg/log"

	"github.com/gin-gonic/gin"
)

var (
	svc  *services.Service
	ctrl *controllers.Controller
)

// New 创建 HTTP 服务
func New(cfg *conf.Config) error {

	// services register
	svc = services.New(cfg)
	// controllers register
	ctrl = controllers.New(svc)

	gin.SetMode(cfg.RunMode)
	engine := gin.New()

	// middlewares
	engine.Use(middlewares.Logger())
	engine.Use(middlewares.CORSHandler())
	engine.Use(middlewares.ConsumeUser())

	// register rest router
	innerRest(engine)

	// register graphql router
	innerGraphql(engine)

	srv := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	log.Infof(context.Background(), "Starting Server %s\n", srv)

	engine.NoRoute(notFound)
	engine.NoMethod()

	err := engine.Run(srv)
	return err
}

func notFound(c *gin.Context) {
	notFound := `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Not Found</title>
  <style>
    html, body {
      background-color: #fff;
      color: #636b6f;
      font-family: 'Nunito', sans-serif;
      font-weight: 100;
      height: 100vh;
      margin: 0;
    }

    .full-height {
      height: 100vh;
    }

    .flex-center {
      align-items: center;
      display: flex;
      justify-content: center;
    }

    .position-ref {
      position: relative;
    }

    .code {
      border-right: 2px solid;
      font-size: 26px;
      padding: 0 15px 0 15px;
      text-align: center;
    }

    .message {
      font-size: 18px;
      text-align: center;
    }
  </style>
</head>
<body>
<div class="flex-center position-ref full-height">
  <div class="code">404</div>
  <div class="message" style="padding: 10px;">Not Found</div>
</div>
</body>
</html>
`
	c.Writer.WriteHeader(http.StatusNotFound)
	_, _ = c.Writer.Write([]byte(notFound))
}

// func serverError(c *gin.Context) {
//	c.String(http.StatusInternalServerError, "Server Error")
// }
