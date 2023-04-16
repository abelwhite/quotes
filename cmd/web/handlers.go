// File Name: handlers.go
package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

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

	http.Redirect(w, r, "/quote/show", http.StatusSeeOther)

}

func (app *application) quoteShow(w http.ResponseWriter, r *http.Request) {

	q, err := app.quote.Read()
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	//an instance of templateData
	data := &templateData{
		Quote: q,
	} //this allows us to pass in mutliple data into the template

	//display quotes using tmpl
	ts, err := template.ParseFiles("./static/html/quoteshow.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	//auming no error
	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}

}

func (app *application) quoteDelete(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve the quote ID from the URL query parameters
	quoteIDStr := r.URL.Query().Get("quote_id")

	// Convert the quote ID string to an integer
	quoteID, err := strconv.Atoi(quoteIDStr)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	// Call the Delete method to remove the quote from the database
	err = app.quote.Delete(quoteID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect the user back to the quote list page
	http.Redirect(w, r, "/quote/show", http.StatusSeeOther)
}

func (app *application) quoteUpdate(w http.ResponseWriter, r *http.Request) {
	// //parse the quote ID from the URL path
	// params := mux.Vars(r)
	// quoteID, err := strconv.Atoi(params["id"])
	// if err != nil {
	// 	http.Error(w,
	// 		http.StatusText(http.StatusBadRequest),
	// 		http.StatusBadRequest)
	// 	return
	// }

	// //get the updated quote data from the request body
	// q := &Quote{}
	// err = json.NewDecoder(r.Body).Decode(q)
	// if err != nil {
	// 	http.Error(w,
	// 		http.StatusText(http.StatusBadRequest),
	// 		http.StatusBadRequest)
	// 	return
	// }

	// //set the QuoteID field of the quote to the parsed quote ID
	// q.QuoteID = quoteID

	// //update the quote in the database
	// err = app.quote.Update(q)
	// if err != nil {
	// 	http.Error(w,
	// 		http.StatusText(http.StatusInternalServerError),
	// 		http.StatusInternalServerError)
	// 	return
	// }

	// //return a success status
	// w.WriteHeader(http.StatusOK)

}
