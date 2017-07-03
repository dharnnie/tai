package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// ServeResource serves static files(css and js) from server
func ServeResource(w http.ResponseWriter, r *http.Request) {
	path := "templates/" + r.URL.Path
	http.ServeFile(w, r, path)
}

// Home Serves the Landing page
func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index-slider.html")
	if err != nil {
		log.Println("Could not parse Home tpl")
		return
	}
	t.Execute(w, nil)
}
