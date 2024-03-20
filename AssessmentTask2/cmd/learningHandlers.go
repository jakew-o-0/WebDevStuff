package main

import (
	"fmt"
	"net/http"
)


func (a *Application) learningContentGet(w http.ResponseWriter, r *http.Request) {
    page,err := getPageParams(w,r)
    if err != nil {
	serverError(w, err.Error())
    	return 
    }
    
    content,err := a.DB.PageLearngingContent(r.Context(), int64(page*10)) 
    if err != nil {
	serverError(w, err.Error())
    	return 
    }

    p := a.newPage(r.Context(), nil)
    p.LearningPage(content, page+1).Render(r.Context(), w)
}

func (a *Application) learingContentPost(w http.ResponseWriter, r *http.Request) {
    searchTerm := r.FormValue("search")
    searchTerm = fmt.Sprintf("%%%s%%", searchTerm)

    content,err := a.DB.GetLearingContent(r.Context(), searchTerm)
    if err != nil {
	serverError(w, fmt.Sprintf(
	    "Failed to query database. error=%q",
	    err.Error()))
    	return 
    }

    p := a.newPage(r.Context(), nil)
    p.LearngingContentComp(content).Render(r.Context(), w)
}
