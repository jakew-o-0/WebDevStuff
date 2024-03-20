package main

import (
	"net/http"
	"t2/templates"
)

func (a *Application) indexPageGet(w http.ResponseWriter, r *http.Request) {
    p := a.newPage(r.Context(), nil)
    p.IndexPage().Render(r.Context(), w)
}


func (a *Application) attractionsGet(w http.ResponseWriter, r *http.Request) {
    page, err := getPageParams(w,r)
    if err != nil {
	serverError(w, err.Error())
	return
    }

    attractions,err := a.DB.PageAttractions(r.Context(), int64(page*10))
    if err != nil {
	serverError(w, err.Error())
    	return 
    }
    
    params := templates.AttractionsParams {
	Attractions: attractions,
	Page: page + 1,
    }
    p := a.newPage(r.Context(), nil)

    if page > 0 {
	p.AttractionsComp(params).Render(r.Context(), w)
    } else {
	p.AttractionsPage(params).Render(r.Context(), w)
    }
}


func (a *Application) facilitiesGet(w http.ResponseWriter, r *http.Request) {
    page, err := getPageParams(w,r)
    if err != nil {
	serverError(w, err.Error())
	return
    }
    

    facilities,err := a.DB.PageFacilities(r.Context(), int64(page*10))
    if err != nil {
	serverError(w, err.Error())
    	return 
    }
    
    params := templates.FacilitiesParams {
	Attractions: facilities,
	Page: page + 1,
    }
    p := a.newPage(r.Context(), nil)

    if page > 0 {
	p.FacilitiesComp(params).Render(r.Context(), w)
    } else {
	p.FacilitiesPage(params).Render(r.Context(), w)
    }
}
