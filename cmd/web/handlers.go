// File Name: handlers.go
package main

import (
	"log"
	"net/http"

	"github.com/abelwhite/quotes/helpers"
)

// creating handler function called greeting
// handler is called when we hit an end point
func (app *application) quoteCreateShow(w http.ResponseWriter, r *http.Request) {
	helpers.RenderTemplates(w, "./static/html/quote.page.tmpl")

}

func (app *application) quoteCreateSubmit(w http.ResponseWriter, r *http.Request) {

	//get the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	quote := r.PostForm.Get("quote") //insert question into the database
	author := r.PostForm.Get("author_name")
	log.Printf("%s %s\n", quote, author)
	id, err := app.quote.Insert(quote, author)
	log.Printf("%s %s %d\n", quote, author, id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

}
