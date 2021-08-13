package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"github.com/favecode/note-core/database"
	"github.com/favecode/note-core/graph/generated"
	graph "github.com/favecode/note-core/graph/resolver"
	customMiddleware "github.com/favecode/note-core/middleware"
	"github.com/favecode/note-core/service"
)

const defaultPort = "8080"

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
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
		Debug:            true,
	}).Handler)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(customMiddleware.AuthMiddleware(user))

	s := service.New(service.Service{
		User:     user,
		Password: database.Password{DB: DB},
	})

	c := generated.Config{Resolvers: &graph.Resolver{Service: s}}
	queryHandler := handler.GraphQL(generated.NewExecutableSchema(c))

	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", queryHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
