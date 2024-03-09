package server

import (
	"database/sql"
	_ "embed"
	"net/http"
	"time"
	"website/cmd/databse"

	"github.com/alexedwards/scs/sqlite3store"
	"github.com/alexedwards/scs/v2"
	_ "github.com/mattn/go-sqlite3"
)


type Server struct{
    DBConn *sql.DB
    database *database.Queries
    SessionManager *scs.SessionManager
}



func New(dbSchema *string) (Server, error) {
    db,err := createDatabaseConn(dbSchema)
    if err != nil {
    	return Server{}, err
    }

    // create server object
    server := Server {
	DBConn: db,
	database: database.New(db),
	SessionManager: createSessionManager(db),
    }
    return server,nil
}


func createDatabaseConn(dbSchema *string) (*sql.DB, error) {
    // create connection
    db, err := sql.Open("sqlite3", "./cmd/databse/db.sqlite3")
    if err != nil {
    	return nil,err
    }

    // create database based off of the schema
    if _,err := db.Exec(*dbSchema); err != nil {
	return nil,err
    }

    return db,nil
}


func createSessionManager(db *sql.DB) *scs.SessionManager {
    sessionManager := scs.New()
    sessionManager.Lifetime = 24 * time.Hour
    sessionManager.Cookie.HttpOnly = true
    sessionManager.Cookie.Path = "/"
    sessionManager.Cookie.SameSite = http.SameSiteLaxMode
    sessionManager.Cookie.Secure = true

    sessionManager.Store = sqlite3store.New(db)

    return sessionManager
}
