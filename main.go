package main

import (
	"Atomic/ent"
	"Atomic/graph"
	"Atomic/graph/generated"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"
const POSTGRES_URL = "postgres://agemnhsp:XZZDblKaDCrB8UShQC4Q8ombv-Y_WLlo@salt.db.elephantsql.com/agemnhsp"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	client := newClient()
	defer client.Close()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Client: client,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func newClient() *ent.Client {
	client, err := ent.Open("postgres", POSTGRES_URL)
	if err != nil {
		log.Fatalf("failed opening connection to postgres : %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
