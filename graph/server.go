package graph

import (
	"fmt"
	"go-graphql-blog/graph/database"
	"go-graphql-blog/graph/generated"
	"go-graphql-blog/graph/utils"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{}}))

	err := database.Connect(utils.GetValue("DATABASE_NAME"))
	if err != nil {
		log.Fatalf("Cannot connect to the database: %v\n", err)
	}

	fmt.Println("Connected to the database")

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
