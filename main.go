package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"github.com/favecode/plog-core/config"
	"github.com/favecode/plog-core/database"
	"github.com/favecode/plog-core/graph/generated"
	graph "github.com/favecode/plog-core/graph/resolver"
	customMiddleware "github.com/favecode/plog-core/middleware"
	"github.com/favecode/plog-core/service"
	"github.com/favecode/plog-core/util"
)

const defaultPort = "8080"

func init() {
	godotenv.Load()
}

func main() {
	DB := database.New()

	defer DB.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := chi.NewRouter()

	user := database.User{DB: DB}

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"X-PINGOTHER", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	// router.Use(customMiddleware.AuthMiddleware(user))

	s := service.New(service.Service{
		User:         user,
		Password:     database.Password{DB: DB},
		Post:         database.Post{DB: DB},
		Connection:   database.Connection{DB: DB},
		Tag:          database.Tag{DB: DB},
		PostTag:      database.PostTag{DB: DB},
		PostLike:     database.PostLike{DB: DB},
		PostSave:     database.PostSave{DB: DB},
		Comment:      database.Comment{DB: DB},
		CommentLike:  database.CommentLike{DB: DB},
		OnlineUser:   database.OnlineUser{DB: DB},
		Notification: database.Notification{DB: DB},
	})

	c := generated.Config{Resolvers: &graph.Resolver{Service: s}}
	queryHandler := handler.New(generated.NewExecutableSchema(c))
	queryHandler.AddTransport(transport.POST{})
	queryHandler.AddTransport(transport.Websocket{
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			claims := jwt.MapClaims{}

			util.ParseJWTWithClaims(initPayload.Authorization(), &claims)

			user, err := user.GetUserByID(claims["jti"].(string))
			if err != nil {
				return nil, err
			}

			currentOnlineUser := &customMiddleware.OnlineUserContext{
				User:      *user,
				Token:     initPayload.Authorization(),
				SocketID:  util.RandomString(20),
				UserAgent: "UserAgent",
			}

			c := context.WithValue(ctx, config.CURRENT_ONLINE_USER_KEY, currentOnlineUser)

			s.AddOnlineUser(c, currentOnlineUser)

			return c, nil
		},
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	queryHandler.Use(extension.Introspection{})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", queryHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
