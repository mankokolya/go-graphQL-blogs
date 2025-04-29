package main

import (
	"fmt"
	"go-graphql-blog/graph"
	"go-graphql-blog/graph/database"
	"go-graphql-blog/graph/generated"
	"go-graphql-blog/graph/middleware"
	"go-graphql-blog/graph/utils"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
)

const defaultPort = "8080"

func NewGraphQLHandler() *chi.Mux {
	var router *chi.Mux = chi.NewRouter()

	router.Use(middleware.NewMiddleware())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	return router
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	var handler *chi.Mux = NewGraphQLHandler()

	err := database.Connect(utils.GetValue("DATABASE_NAME"))
	if err != nil {
		log.Fatalf("Cannot connect to the database: %v\n", err)
	}

	fmt.Println("Connected to the database")

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
