package server

import (
	"go-starter/internal/graphql/generated"
	graph "go-starter/internal/graphql/resolvers"
	"net/http"

	"strings"
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

// newGraphQLServer creates a GraphQL server.
func newGraphQLServer(es graphql.ExecutableSchema) *handler.Server {
	srv := handler.New(es)
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

	return srv
}

// graphqlHandler defines the GraphQL handler.
func graphqlHandler() gin.HandlerFunc {

	config := generated.Config{
		Resolvers: &graph.Resolver{
			Svc: svc,
		},
	}

	srv := newGraphQLServer(generated.NewExecutableSchema(config))

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

// playgroundHandler defines the Playground handler.
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// registerGraphqlRouter registers the GraphQL router.
func registerGraphqlRouter(e *gin.Engine, mode string) {
	// GraphQL
	g := e.Group("/graphql")
	g.POST("", graphqlHandler())
	if !strings.Contains("release", mode) {
		g.GET("", playgroundHandler())
	}
}
