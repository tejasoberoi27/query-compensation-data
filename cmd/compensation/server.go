package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"query-compensation-data/internal/compensation/config"
	"query-compensation-data/internal/compensation/env"
	"query-compensation-data/internal/compensation/handler/graph/compensation/generated"
	compensation_service "query-compensation-data/internal/compensation/service/compensation"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	compensation_handler "query-compensation-data/internal/compensation/handler/compensation"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	router := chi.NewRouter()
	port := cfg.APIServer.Port
	env, err := env.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create environment: %v", err)
	}

	svc := compensation_service.New(env)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{Resolvers: compensation_handler.NewResolver(svc)},
	))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/comp_data/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
