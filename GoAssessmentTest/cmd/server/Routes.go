package server

import (
	"net/http"
)

func (s *Server) IndexPageGet(w http.ResponseWriter, r *http.Request) {
    // create page structure and render it
    page := s.NewPage(r)
    page.BaseIndexPage().Render(r.Context(), w)
}

func (srv *Server) InboxPageGet(w http.ResponseWriter, r *http.Request) {
}
