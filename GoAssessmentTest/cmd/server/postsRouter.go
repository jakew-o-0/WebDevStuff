package server

import (
	"net/http"
	"website/templates"

	"github.com/go-chi/chi/v5"
)

func (serv *Server) PostsRoutes() chi.Router {
    router := chi.NewRouter()
    router.Get("/", serv.PostsPageGet)

    return router
}



func (srv *Server) PostsPageGet(w http.ResponseWriter, r *http.Request) {
    page := templates.Base{
	IsLoggedIn: false,
	Child: templates.PostsPage{},
    }
    page.RenderPage().Render(r.Context(), w)
}


func (s *Server) PostsCreateGet(w http.ResponseWriter, r *http.Request) {
    page := templates.Base{
	IsLoggedIn: true,
	Child: templates.CreatePostModal{},
    }
    page.RenderPage().Render(r.Context(), w)
}
