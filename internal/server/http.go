package server

import (
	"github.com/gin-gonic/gin"
	"go-starter/common/conf"
	"go-starter/common/types"
	"go-starter/internal/controller"
	"go-starter/internal/middleware"
	"go-starter/internal/service"
	"net/http"
)

var (
	svc  *service.Service
	ctrl *controller.Controller
)

// New 创建 HTTP 服务
func New(cfg *conf.Config) (*gin.Engine, *service.Service) {

	// service register
	svc = service.New(cfg)
	// controller register
	ctrl = controller.New(svc)

	gin.SetMode(cfg.RunMode)
	engine := gin.New()

	// middleware
	engine.Use(middleware.Logger())
	engine.Use(middleware.CORSHandler())
	engine.Use(middleware.ConsumeUser())

	// register rest router
	innerRest(engine)

	// register graphql router
	innerGraphql(engine, cfg.RunMode)

	// srv := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	// log.Infof(context.Background(), "Starting Server %s\n", srv)

	engine.NoRoute(notFound)
	engine.NoMethod()

	// err := engine.Run(srv)
	return engine, svc
}

func notFound(c *gin.Context) {
	// 	notFound := `<!DOCTYPE html>
	// <html lang="en">
	// <head>
	//   <meta charset="utf-8">
	//   <meta name="viewport" content="width=device-width, initial-scale=1">
	//   <title>Not Found</title>
	//   <style>
	//     html, body {
	//       background-color: #fff;
	//       color: #636b6f;
	//       font-family: 'Nunito', sans-serif;
	//       font-weight: 100;
	//       height: 100vh;
	//       margin: 0;
	//     }
	//
	//     .full-height {
	//       height: 100vh;
	//     }
	//
	//     .flex-center {
	//       align-items: center;
	//       display: flex;
	//       justify-content: center;
	//     }
	//
	//     .position-ref {
	//       position: relative;
	//     }
	//
	//     .code {
	//       border-right: 2px solid;
	//       font-size: 26px;
	//       padding: 0 15px 0 15px;
	//       text-align: center;
	//     }
	//
	//     .message {
	//       font-size: 18px;
	//       text-align: center;
	//     }
	//   </style>
	// </head>
	// <body>
	// <div class="flex-center position-ref full-height">
	//   <div class="code">404</div>
	//   <div class="message" style="padding: 10px;">Not Found</div>
	// </div>
	// </body>
	// </html>
	// `
	// 	c.Writer.WriteHeader(http.StatusNotFound)
	// 	_, _ = c.Writer.Write([]byte(notFound))

	c.JSON(http.StatusNotFound, types.JSON{"message": "Not Found"})
}
