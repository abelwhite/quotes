// filename: routes.go
// we are using a 3rd party router
// router provides an end poin that the user type in.
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// create a multiplexer
	router := httprouter.New() //httprouter from juienschmidft/httprouter
	// create a file server
	// filer server needs to be created to show static content
	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer)) //remove "resources"

	router.HandlerFunc(http.MethodGet, "/quote/create", app.quoteCreateShow)    //provide string and hander
	router.HandlerFunc(http.MethodPost, "/quote/create", app.quoteCreateSubmit) //provide string and hander

	return router
} //router is the data structure to allow us to locate the end points
