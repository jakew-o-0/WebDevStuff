package main

import (
	"database/sql"
	_ "embed"
	"log"
	"net/http"
	"t2/internal/database"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_"github.com/mattn/go-sqlite3"
)


// shares values accross route handlers
// route handlers are attached to this struct
// dependancy injection pattern
type Application struct {
    DB *database.Queries
    Sessions *scs.SessionManager
}


func main() {
    // create db connection
    // propagate error
    // close db when function returns
    db, err := sql.Open("sqlite3", "./internal/database/db.sqlite3")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()


    a := newApplication(db)
    r := chi.NewRouter()


    // middleware
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Use(a.Sessions.LoadAndSave)

    // routes
    r.Get("/", a.indexPageGet)

    r.Get("/auth/login", a.loginPageGet)
    r.Post("/auth/login", a.loginPagePost)
    r.Get("/auth/signup", a.signupPageGet)
    r.Post("/auth/signup", a.signupPagePost)
    r.Post("/auth/logout", a.logoutPost)

    r.Get("/attractions/{page}", a.attractionsGet)
    r.Get("/facilities/{page}", a.facilitiesGet)
    r.Get("/learning/{page}", a.learningContentGet)
    r.Post("/learning", a.learingContentPost)

    // start server
    log.Fatal(http.ListenAndServe(":3000", r))
}
