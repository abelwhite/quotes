package helpers

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplates(w http.ResponseWriter, tmpl string) {
	//see a webpage display
	//template->to write html and pass dynamic data
	ts, err := template.ParseFiles(tmpl)
	if err != nil { //parse file did not do its job
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil) //nill because there is no data to inject
	if err != nil {          //parse file did not do its job
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
