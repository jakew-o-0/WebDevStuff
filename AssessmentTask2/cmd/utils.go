package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"t2/internal/database"
	"t2/templates"

	"github.com/alexedwards/scs/sqlite3store"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
)

func (a *Application) newPage(ctx context.Context, errors map[string]string) templates.Page {
    return templates.Page{
	LoggedIn: a.isLoggedIn(ctx),
	Errors: errors,
	Flash: a.Sessions.PopString(ctx, "FlashMessage"),
    }
}

func (a *Application)isLoggedIn(ctx context.Context) bool {
    _,ok := a.Sessions.Get(ctx, "AuthenticatedUserID").(int64)
    if !ok {
    	return false
    }
    return true 
}


func newApplication(db *sql.DB) Application {
    // create database based off of schema
    err := database.CreateFromSchema(db)
    if err != nil {
	panic(err.Error())
    }

    // create session manager
    s := scs.New()
    s.Store = sqlite3store.New(db)

    return Application{
	DB: database.New(db),
	Sessions: s,
    }
}


func serverError(w http.ResponseWriter, errMsg string) {
    http.Error(w, errMsg ,http.StatusInternalServerError)
}

func getPageParams(w http.ResponseWriter, r *http.Request) (int,error) {
    pageStr := chi.URLParam(r, "page")
    page,err := strconv.Atoi(pageStr)
    if err != nil {
	err = errors.New(fmt.Sprintf("Invalid Page Param. error=%q", err.Error()))
	return 0,err
    }
    return page,nil
}








func (a *Application) createAttractions() {
    attractions := database.AttractionDummyData()
    for _,att := range attractions {
	params := database.CreateAttractionParams{
	    Title: att.Title,
	    Content: att.Content,
	}
	err := a.DB.CreateAttraction(context.Background(), params)
	if err != nil {
	    panic("failed to create attractions")
	}
    }
}

func (a *Application) createFacilities() {
    f := database.FacilityDummyData()
    for _,ff := range f {
	params := database.CreateFacilityParams {
	    Title: ff.Title,
	    Content: ff.Content,
	}
	err := a.DB.CreateFacility(context.Background(), params)
	if err != nil {
	    panic("failed to create facilities")
	}
    }
}

func (a *Application) createFactfile() {
    f := database.CreateDummyFactFiles()
    for _,ff := range f {
	params := database.CreateFactFileParams {
	    Name: ff.Name,
	    Description: ff.Description,
	    Habitat: ff.Habitat,
	    Diet: ff.Diet,
	    ConservationStatus: ff.ConservationStatus,
	}
	err := a.DB.CreateFactFile(context.Background(), params)
	if err != nil {
	    panic(err.Error())
	}
    }
}
