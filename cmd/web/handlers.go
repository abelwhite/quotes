// File Name: handlers.go
package main

import (
	"net/http"
)

// creating handler function called greeting
// handler is called when we hit an end point
func (app *application) quoteCreateShow(w http.ResponseWriter, r *http.Request) {
	// helpers.RenderTemplates(w, "./static/html/poll.page.tmpl")

}

func (app *application) quoteCreateSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		//set header
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//get the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	the_question := r.PostForm.Get("new_question") //insert question into the database
	_, err = app.question.Insert(the_question)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

}
