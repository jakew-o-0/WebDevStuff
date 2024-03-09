package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"time"
	"website/cmd/server"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//go:embed database/schema.sql
var dbSchema string

func main() {
    r := chi.NewRouter()

    // create server object
    server,err := server.New(&dbSchema)
    if err != nil {
	fmt.Printf("Failed to create Database: %v\n", err.Error())
    	panic(0)
    }
    defer server.DBConn.Close()


    // setup middleware
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Use(middleware.Timeout(1 * time.Minute))
    r.Use(server.SessionManager.LoadAndSave)

    // setup http routes
    r.Get("/", server.IndexPageGet)
    r.Mount("/posts", server.PostsRoutes())
    r.Mount("/auth", server.AuthRoutes())

    // launch server
    log.Fatal(http.ListenAndServe(":3000", r))
}
