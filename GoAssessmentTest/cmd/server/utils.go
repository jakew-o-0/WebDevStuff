package server

import (
	"fmt"
	"net/http"
	"website/templates"
)

func (s *Server) NewPage(r *http.Request) templates.Base{
    return templates.Base {
	FlashMessage: s.SessionManager.PopString(r.Context(), "FlashMessage"),
	LoggedIn: s.SessionManager.Exists(r.Context(), "AuthenticatedUserID"),
	Errors: nil,
    }
}

func ServerError(w http.ResponseWriter, r *http.Request, err error) {
    errorMsg := fmt.Sprintf("error=%v", err.Error())
    http.Error(w, errorMsg, http.StatusInternalServerError)
}

func UnauthedError(w http.ResponseWriter, r *http.Request, err error) {
    errorMsg := fmt.Sprintf("error=%v", err.Error())
    http.Error(w, errorMsg, http.StatusUnauthorized)
}


