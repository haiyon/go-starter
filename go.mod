module go-starter

go 1.16

require (
	entgo.io/ent v0.9.0
	github.com/99designs/gqlgen v0.13.0
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.7.3
	github.com/go-sql-driver/mysql v1.5.1-0.20200311113236-681ffa848bae
	github.com/gorilla/websocket v1.4.2
	github.com/jackc/pgx/v4 v4.13.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/swaggo/swag v1.7.0
	github.com/vektah/gqlparser/v2 v2.1.0
	go-starter/common v0.0.0-00010101000000-000000000000
)

replace go-starter/common => ./common
