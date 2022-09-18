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

	"github.com/plogto/core/constants"
	"github.com/plogto/core/database"
	"github.com/plogto/core/graph/generated"
	graph "github.com/plogto/core/graph/resolver"
	customMiddleware "github.com/plogto/core/middleware"
	"github.com/plogto/core/service"
	"github.com/plogto/core/util"
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

	users := database.Users{DB: DB}

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"X-PINGOTHER", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(customMiddleware.AuthMiddleware(users))

	s := service.New(service.Service{
		Users:                                 users,
		Passwords:                             database.Passwords{DB: DB},
		Posts:                                 database.Posts{DB: DB},
		Files:                                 database.Files{DB: DB},
		Connections:                           database.Connections{DB: DB},
		CreditTransactions:                    database.CreditTransactions{DB: DB},
		CreditTransactionInfos:                database.CreditTransactionInfos{DB: DB},
		CreditTransactionTemplates:            database.CreditTransactionTemplates{DB: DB},
		CreditTransactionDescriptionVariables: database.CreditTransactionDescriptionVariables{DB: DB},
		Tickets:                               database.Tickets{DB: DB},
		TicketMessages:                        database.TicketMessages{DB: DB},
		Tags:                                  database.Tags{DB: DB},
		TicketMessageAttachments:              database.TicketMessageAttachments{DB: DB},
		PostAttachments:                       database.PostAttachments{DB: DB},
		PostTags:                              database.PostTags{DB: DB},
		LikedPosts:                            database.LikedPosts{DB: DB},
		SavedPosts:                            database.SavedPosts{DB: DB},
		InvitedUsers:                          database.InvitedUsers{DB: DB},
		OnlineUsers:                           database.OnlineUsers{DB: DB},
		Notifications:                         database.Notifications{DB: DB},
		NotificationTypes:                     database.NotificationTypes{DB: DB},
	})

	s.OnlineUsers.DeleteAllOnlineUsers()

	c := generated.Config{Resolvers: &graph.Resolver{Service: s}}
	queryHandler := handler.New(generated.NewExecutableSchema(c))
	queryHandler.AddTransport(transport.POST{})
	queryHandler.AddTransport(transport.MultipartForm{
		MaxMemory:     32 * constants.MB,
		MaxUploadSize: 50 * constants.MB,
	})
	queryHandler.AddTransport(transport.Websocket{
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			claims := jwt.MapClaims{}

			token, _ := customMiddleware.StripBearerPrefixFromToken(initPayload.Authorization())
			util.ParseJWTWithClaims(token, &claims)

			if len(claims) > 0 {
				user, err := users.GetUserByID(claims["jti"].(string))

				if err != nil {
					return nil, err
				}

				currentOnlineUser := &customMiddleware.OnlineUserContext{
					User:      *user,
					Token:     token,
					SocketID:  util.RandomString(20),
					UserAgent: "UserAgent",
				}

				c := context.WithValue(ctx, constants.CURRENT_ONLINE_USER_KEY, currentOnlineUser)

				s.AddOnlineUser(c, currentOnlineUser)

				return c, nil
			}

			return ctx, nil
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
