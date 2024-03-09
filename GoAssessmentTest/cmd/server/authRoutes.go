package server

import (
	"database/sql"
	"net/http"
	database "website/cmd/databse"
	"website/cmd/validator"
	"website/templates"

	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)


type loginForm struct {
    Username string
    Password string
}

type signupForm struct {
    username string
    password string
    confPassword string
}


func (s *Server) AuthRoutes() chi.Router {
    r := chi.NewRouter()
    r.Get("/login", s.loginGet)
    r.Post("/login", s.loginPost)
    r.Get("/signup", s.signupGet)
    r.Post("/signup", s.signupPost)
    r.Post("/logout", s.logoutPost)
    r.Delete("/", s.deleteUser)

    return r
}



func (s *Server) loginGet(w http.ResponseWriter, r *http.Request) {
    page := s.NewPage(r)
    page.BaseLoginPage().Render(r.Context(), w)
}

func (s *Server) loginPost(w http.ResponseWriter, r *http.Request) {
    // get form values
    loginForm := loginForm {
	Username: r.FormValue("username"),
	Password: r.FormValue("password"),
    }

    // validate
    validator := s.validateLoginForm(w,r, loginForm)
    if !validator.IsValid() { // on error send a login form with error messages
	page := s.NewPage(r)
	page.Errors = validator.Errors
	page.LoginPage().Render(r.Context(), w)
	return
    }

    // make session an authenticated one
    s.SessionManager.RenewToken(r.Context())
    s.SessionManager.Put(r.Context(), "AuthenticatedUserID", nil)
    s.SessionManager.Put(r.Context(), "FlashMessage", "Successfully Logged In")

    // redirect to posts
    templates.HtmxRedirect("/posts").Render(r.Context(), w)
}






func (s *Server) signupGet(w http.ResponseWriter, r *http.Request) {
    page := s.NewPage(r)
    page.BaseSignupPage().Render(r.Context(), w)
}

func (s *Server) signupPost(w http.ResponseWriter, r *http.Request) {
    // get form values
    signupForm := signupForm {
	username: r.FormValue("username"),
	password: r.FormValue("password"),
	confPassword: r.FormValue("confPassword"),
    }

    // validate
    validator := s.validateSignupForm(signupForm)
    if !validator.IsValid() {
	page := s.NewPage(r)
	page.Errors = validator.Errors
	page.SignupPage().Render(r.Context(), w)
	return
    }

    // create user in database
    hashedPassword,err := bcrypt.GenerateFromPassword([]byte(signupForm.password), 12)
    if err != nil {
	ServerError(w,r,err)
    	return 
    }
    
    user := database.CreateUserParams {
	Username: signupForm.username,
	Password: string(hashedPassword),
    }
    userID,err := s.database.CreateUser(r.Context(), user)
    if err != nil {
	ServerError(w,r,err)
    	return 
    }

    // create session
    s.SessionManager.RenewToken(r.Context())
    s.SessionManager.Put(r.Context(), "AuthenticatedUserID", userID)
    s.SessionManager.Put(r.Context(), "FlashMessage", "Successfully Created Account")

    // redirect to posts
    templates.HtmxRedirect("/posts").Render(r.Context(), w)
}


func (s *Server) logoutPost(w http.ResponseWriter, r *http.Request) {
    s.SessionManager.RenewToken(r.Context())
    s.SessionManager.Remove(r.Context(), "AuthenticatedUserID")
    s.SessionManager.Put(r.Context(), "FlashMessage", "Successfully Logged Out")

    templates.HtmxRedirect("/").Render(r.Context(), w)
}


func (s *Server) deleteUser(w http.ResponseWriter, r *http.Request) {
}













func (s *Server) validateSignupForm(f signupForm) validator.Validator {
    validator := validator.New()
    validator.Validate(
	validator.ValidateMinMax(f.username, 2,24),
	"Username",
	"Invalid Username length: Must be between 2 and 24 chars")

    validator.Validate(
	validator.ValidateUsername(f.username),
	"Username",
	"Invalid Username format: Must only have letters and numbers")

    validator.Validate(
	validator.ValidateStrings(f.password, f.confPassword),
	"Password",
	"Passwords do not match")

    validator.Validate(
	validator.ValidatePassword(f.password),
	"Password",
	"password must have 1 uppercase and lowercase letters, 1 number and 1 special character")

    validator.Validate(
	validator.ValidateMinMax(f.password, 8,32),
	"Password",
	"password must be between 8 and 32 chars long")

    return validator
}


func (s *Server) validateLoginForm(w http.ResponseWriter, r *http.Request, l loginForm) validator.Validator {
    validator := validator.New()
    validator.Validate(
	validator.ValidateMinMax(l.Username, 2,24),
	"Username",
	"Invalid Username or Password")

    validator.Validate(
	validator.ValidateUsername(l.Username),
	"Username",
	"Invalid Username or Password")

    validator.Validate(
	validator.ValidateMinMax(l.Password, 8,32),
	"Password",
	"Invalid Username or Password")


    // validate username and password based off of querying the database
    userID,err := s.database.GetUserByUsername(r.Context(), l.Username)
    if err != nil {
    	if err != sql.ErrNoRows{
	    ServerError(w,r,err)
	}

	validator.Validate(
	    false,
	    "Username",
	    "Invalid Username or Password")

	return validator
    }

    err = bcrypt.CompareHashAndPassword([]byte(userID.Password), []byte(l.Password))
    if err != nil {
	validator.Validate(
	    false,
	    "Password",
	    "Invalid Username or Password")
    }
    
    return validator
}
