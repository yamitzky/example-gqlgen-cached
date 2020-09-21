package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/yamitzky/gqlgentest/graph"
	"github.com/yamitzky/gqlgentest/graph/expiration"
	"github.com/yamitzky/gqlgentest/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	cache := graph.NewQueryCache()

	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv.AddTransport(expiration.TransportGET{})
	srv.AddTransport(transport.POST{})
	srv.Use(extension.AutomaticPersistedQuery{Cache: cache})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", expiration.Middleware(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
