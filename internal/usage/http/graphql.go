package http

import (
	"net/http"
	generated "haiyon/go-starter/internal/generated/graphql"
	"haiyon/go-starter/internal/resolver"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// newGraphQLServer 创建 GraphQL 服务
func newGraphQLServer(es graphql.ExecutableSchema) (srv *handler.Server) {
	srv = handler.New(es)
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		KeepAlivePingInterval: 15 * time.Second,
	})

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	return
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {

	config := generated.Config{
		Resolvers: &resolver.Resolver{
			Svc: svc,
		},
	}

	srv := newGraphQLServer(generated.NewExecutableSchema(config))

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func innerGraphql(e *gin.Engine) {
	// GraphQL
	g := e.Group("/graphql")
	g.POST("", graphqlHandler())
	g.GET("/playground", playgroundHandler())
}
