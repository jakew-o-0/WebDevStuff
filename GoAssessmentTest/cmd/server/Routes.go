package server

import "net/http"

func (srv *Server) IndexPageGet(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello"))
}
