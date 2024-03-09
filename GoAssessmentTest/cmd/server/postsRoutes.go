package server

import (
	"errors"
	"net/http"
	"strconv"
	database "website/cmd/databse"
	"website/cmd/validator"
	"website/templates"

	"github.com/go-chi/chi/v5"
)

func (s *Server) PostsRoutes() chi.Router {
    router := chi.NewRouter()
    router.Get("/", s.PostsPageGet)
    router.Post("/", s.PostsPagePost)
    router.Get("/create", s.PostsCreateGet)
    return router
}



func (s *Server) PostsPageGet(w http.ResponseWriter, r *http.Request) {
    // get page query params
    pageParam := r.URL.Query().Get("page")
    if pageParam == "" {
	pageParam = "0"
    }
    // convert to an int and propagate error
    pageInt,err := strconv.Atoi(pageParam)
    if err != nil {
	ServerError(w,r,err)
	return
    }
    
    // get get current pages posts
    posts,err := s.database.GetPosts(r.Context(), int64(pageInt*10))
    if err != nil {
	ServerError(w,r,err)
    	return 
    }


    page := s.NewPage(r)
    params := templates.PostsPageParams{
	CurPage: pageInt+1,
	Posts: posts,
    }
    // return just the posts if the page is being paginated
    if pageInt > 0 {
	page.PostsPage(params).Render(r.Context(), w)
	return
    }

    // return the whole page contents
    page.BasePostsPage(params).Render(r.Context(), w)
}


func (s *Server) PostsPagePost(w http.ResponseWriter, r *http.Request) {
    if !s.SessionManager.Exists(r.Context(), "AuthenticatedUserID") {
	err := errors.New("You must be logged in to create a post")
	UnauthedError(w,r,err)
	return
    }

    // get data from form
    r.ParseForm()
    title := r.Form.Get("title")
    content := r.Form.Get("content")


    // validate form data
    validator := validator.New()

    validator.Validate(
	validator.ValidateMinMax(title, 0, 20),
	"Title",
	"Invalid length: must be between 1 and 20 characters long")

    validator.Validate(
	validator.ValidateMinMax(content, 0, 128),
	"Content",
	"Invalid length: must be between 1 and 128 characters long")

    // check if there where any errors
    if !validator.IsValid() {
	templates.PostsModal(
	    validator.Errors["Title"],
	    validator.Errors["Content"],
	    ).Render(r.Context(), w)
	return
    }


    // insert into database
    newPost := database.CreatePostParams{
	Title: title,
	Content: content,
    }
    err := s.database.CreatePost(r.Context(), newPost)
    if err != nil {
	ServerError(w,r,err)
	return
    }


    // redirect to the postspage to reload content
    templates.HtmxRedirect("/posts").Render(r.Context(), w)
}


// server create post modal
func (s *Server) PostsCreateGet(w http.ResponseWriter, r *http.Request) {
    if !s.SessionManager.Exists(r.Context(), "AuthenticatedUserID") {
	err := errors.New("You must be logged in to create a post")
	UnauthedError(w,r,err)
	return
    }
    templates.PostsModal("","").Render(r.Context(), w)
}
