package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"t2/internal/database"
	"t2/internal/validator"
	"t2/templates"

	"golang.org/x/crypto/bcrypt"
)



func (a *Application) loginPageGet(w http.ResponseWriter, r *http.Request) {
    p := a.newPage(r.Context(), nil) 
    p.LoginPage().Render(r.Context(), w)
}

func (a *Application) loginPagePost(w http.ResponseWriter, r *http.Request) {
    // get form encoded data from request
    email := r.FormValue("email")
    password := r.FormValue("password")


    // validate
    v := validator.New()
    v.Validate(
	v.ValidEmail(email),
	"Err",
	"Invalid Email or Password")

    v.Validate(
	v.ValidPassword(password),
	"Err",
	"Invalid Email or Password")

    // get user from db
    user,err := a.DB.GetUserByEmail(r.Context(), email)
    if err != nil {
	if err != sql.ErrNoRows{ 
	    serverError(w, err.Error())
	    return 
	}
	v.Validate(
	    false,
	    "Err",
	    "Invalid Email or Password")
    }

    v.Validate(
	v.ValidPasswordHash([]byte(user.Password), []byte(password)),
	"Err",
	"Invalid Email or Password")

    if !v.IsValid() {
	p := a.newPage(r.Context(), v.Errors)
	p.SignupPageComp().Render(r.Context(), w)
	return
    }


    // create authed session & flash message
    a.Sessions.RenewToken(r.Context())
    a.Sessions.Put(r.Context(), "AuthenticatedUserID", user.UserID)
    a.Sessions.Put(r.Context(), "FlashMessage", "Successfully Logged In")
    
    // redirect to indexpage
    w.Header().Add("HX-Refresh", "true")
    p := a.newPage(r.Context(), nil)
    p.IndexPage().Render(r.Context(), w)
}




func (a *Application) signupPageGet(w http.ResponseWriter, r *http.Request) {
    p := templates.Page {
	LoggedIn: false,
    }
    p.SignupPage().Render(r.Context(), w)
}

func (a *Application) signupPagePost(w http.ResponseWriter, r *http.Request) {
    // get form encoded data from request
    firstName := r.FormValue("firstName")
    lastName := r.FormValue("lastName")
    email := r.FormValue("email")
    password := r.FormValue("password")
    confPassword := r.FormValue("confPassword")


    // validate
    v := validator.New()

    v.Validate(
	v.MinMaxLen(firstName, 3,24),
	"FirstName",
	"Invalid length: must be between 3 and 24 chars")
    v.Validate(
	v.ValidName(firstName),
	"FirstName",
	"Name must only have letters")

    v.Validate(
	v.MinMaxLen(lastName, 3,24),
	"FirstName",
	"Invalid length: must be between 3 and 24 chars")
    v.Validate(
	v.ValidName(lastName),
	"FirstName",
	"Name must only have letters")

    v.Validate(
	v.ValidEmail(email),
	"Email",
	"Invalid Email")

    // check if email exists 
    _,err := a.DB.GetUserByEmail(r.Context(), email)
    if err != sql.ErrNoRows {
	v.Validate(
	    false,
	    "Email",
	    "Email is already in use")
    } 

    v.Validate(
	v.ValidPassword(password),
	"Password",
	"Password must have 1 upper, lowwer, number and special char")
    v.Validate(
	v.StringsMatch(password, confPassword),
	"Password",
	"Passwords dont match")

    if !v.IsValid() {
	p := a.newPage(r.Context(), v.Errors)
	p.SignupPageComp().Render(r.Context(), w)
	return
    }


    // Create User
    hashedPassword,err := bcrypt.GenerateFromPassword([]byte(password), 12)
    if err != nil {
	serverError(w, err.Error())
	return
    }

    user := database.CreateUserParams {
	FirstName: firstName,
	LastName: lastName,
	Email: email,
	Password: string(hashedPassword),
    }
    if err = a.DB.CreateUser(r.Context(), user); err != nil {
	serverError(w, 
	    fmt.Sprintf("Failed to create user in db. Error=%s", err.Error()))
	return
    }

    userID,err := a.DB.GetUserID(r.Context(), email)
    if err != nil {
	serverError(w, err.Error())
    	return 
    }


    // Create Session
    a.Sessions.RenewToken(r.Context())
    a.Sessions.Put(r.Context(), "AuthenticatedUserID", userID)
    a.Sessions.Put(r.Context(), "FlashMessage", "Account Created. You are now signed in.")


    // redirect to index page
    w.Header().Add("HX-Retarget", "body")
    w.Header().Add("HX-Push-Url", "true")
    p := a.newPage(r.Context(), nil)
    p.IndexPage().Render(r.Context(), w)
}

func (a *Application) logoutPost(w http.ResponseWriter, r *http.Request) {
    a.Sessions.RenewToken(r.Context())
    a.Sessions.Destroy(r.Context())
    a.Sessions.Put(r.Context(), "FlashMessage", "Logged Out")

    w.Header().Add("HX-Push-Url", "/")
    p := a.newPage(r.Context(), nil)
    p.IndexPage().Render(r.Context(), w)
}
