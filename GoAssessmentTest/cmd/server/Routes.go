package server

import (
	"net/http"
	"website/templates"
)

func (srv *Server) IndexPageGet(w http.ResponseWriter, r *http.Request) {
    // create page structure and render it
    page := templates.Base{
	IsLoggedIn: false,
	Child: templates.IndexPage{},
    }
    page.RenderPage().Render(r.Context(), w)
}

func (srv *Server) InboxPageGet(w http.ResponseWriter, r *http.Request) {
    page := templates.Base{
	IsLoggedIn: false,
    }
    page.RenderPage().Render(r.Context(), w)
}
