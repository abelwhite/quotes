// File Name: handlers.go
package main

import (
	"html/template"
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

	http.Redirect(w, r, "/quote/show", http.StatusSeeOther)

}

func (app *application) quoteShow(w http.ResponseWriter, r *http.Request) {

	//create SQL statement
	readQuotes := `
		SELECT *
		FROM quotes
		 
	`
	rows, err := app.quote.DB.Query(readQuotes)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	defer rows.Close()

	var quotes []Quote
	for rows.Next() {
		var q Quote
		err = rows.Scan(&q.QuoteID, &q.Quote, &q.Author, &q.CreatedAt)

		if err != nil {
			log.Println(err.Error())
			http.Error(w,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
		quotes = append(quotes, q) //contain first row
	}
	//check to see if there were erroe generated
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	//print the values that are in the slice
	// for _, quote := range quotes {
	// 	fmt.Fprintf(w, "%v \n", quote)
	// }

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
	err = ts.Execute(w, quotes)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}

}
