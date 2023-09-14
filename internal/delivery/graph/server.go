package graph

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func StartGraphQLServer(mux *http.ServeMux) {
	port := "8998"
	handler := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{}}))

	mux.HandleFunc("/", playground.Handler("GraphQL playground", "/query"))
	mux.HandleFunc("/query", handler.ServeHTTP)

	srv := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("GrapghQL Server Started on Port : ", port)

	<-done
	log.Print("GrapghQL Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("GrapghQL Server Shutdown Failed:%+v", err)
	}

	log.Print("GraphQL Server Exited Properly")
}
