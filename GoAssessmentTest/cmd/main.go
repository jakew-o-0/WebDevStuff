package main

import (
	"net/http"
	"time"
	"website/cmd/server"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
    server := server.New()

    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Use(middleware.Timeout(1 * time.Minute))

    r.Get("/", server.IndexPageGet)

    http.ListenAndServe(":3000", r)
}
