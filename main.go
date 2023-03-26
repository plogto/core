package main

import (
	"context"
	"fmt"
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

	_ "github.com/lib/pq"
	"github.com/plogto/core/constants"
	"github.com/plogto/core/database"
	"github.com/plogto/core/db"
	graphDataloader "github.com/plogto/core/graph/dataloader"
	"github.com/plogto/core/graph/generated"
	graphResolver "github.com/plogto/core/graph/resolver"
	customMiddleware "github.com/plogto/core/middleware"
	"github.com/plogto/core/service"
	"github.com/plogto/core/util"
)

const defaultPort = "8080"

func init() {
	godotenv.Load()
}

func main() {
	// FIXME
	newDB, err := db.Open(os.Getenv("NEW_DATABASE_URL"))

	queries := db.New(newDB)

	if err != nil {
		fmt.Println(err)
	}
	// FIXME
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
		Connections:                           database.Connections{Queries: queries},
		CreditTransactionDescriptionVariables: database.CreditTransactionDescriptionVariables{Queries: queries},
		CreditTransactionTemplates:            database.CreditTransactionTemplates{Queries: queries},
		CreditTransactionInfos:                database.CreditTransactionInfos{Queries: queries},
		CreditTransactions:                    database.CreditTransactions{Queries: queries},
		Files:                                 database.Files{Queries: queries},
		InvitedUsers:                          database.InvitedUsers{Queries: queries},
		LikedPosts:                            database.LikedPosts{Queries: queries},
		NotificationTypes:                     database.NotificationTypes{Queries: queries},
		Notifications:                         database.Notifications{Queries: queries},
		Passwords:                             database.Passwords{Queries: queries},
		PostAttachments:                       database.PostAttachments{Queries: queries},
		PostMentions:                          database.PostMentions{Queries: queries},
		PostTags:                              database.PostTags{Queries: queries},
		Posts:                                 database.Posts{Queries: queries},
		Tags:                                  database.Tags{Queries: queries},
		Users:                                 users,
		Tickets:                               database.Tickets{DB: DB},
		TicketMessages:                        database.TicketMessages{DB: DB},
		TicketMessageAttachments:              database.TicketMessageAttachments{DB: DB},
		SavedPosts:                            database.SavedPosts{DB: DB},
		OnlineUsers:                           database.OnlineUsers{DB: DB},
	})

	c := generated.Config{Resolvers: &graphResolver.Resolver{Service: s}}
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

			// if len(claims) > 0 {
			// 	user, err := users.GetUserByID(claims["jti"].(string))

			// 	if err != nil {
			// 		return nil, err
			// 	}

			// 	currentOnlineUser := &customMiddleware.OnlineUserContext{
			// 		User:      *user,
			// 		Token:     token,
			// 		SocketID:  util.RandomString(20),
			// 		UserAgent: "UserAgent",
			// 	}

			// 	c := context.WithValue(ctx, constants.CURRENT_ONLINE_USER_KEY, currentOnlineUser)

			// 	s.AddOnlineUser(c, currentOnlineUser)

			// 	return c, nil
			// }

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
	router.Handle("/query", graphDataloader.DataloaderMiddleware(DB, queries, queryHandler))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
