package server

import "github.com/go-chi/chi/v5"

type Server struct{}

func New() Server {
    return Server{}
}

func (serv *Server) Routes() chi.Router {
    router := chi.NewRouter()

    router.Get("/", serv.IndexPageGet)
    router.Mount("/posts", serv.PostsRoutes())

    return router
}
