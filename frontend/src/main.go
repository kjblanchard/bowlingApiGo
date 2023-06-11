package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// Template data we will pass into the templates
type TemplateData struct {
	Title string
	Body  string
}

// The cached templates
var templates *template.Template

// Main page handler.
func handler(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{
		Title: "Bowling",
		Body:  "What",
	}
	err := templates.ExecuteTemplate(w, "homepage.html", data)
	if err != nil {
		log.Printf("Something happened:\n %s", err.Error())
		fmt.Fprintf(w, "Borked")
	}

	// loadedTemplate.Execute(w, data)
}

// Load all templates from disk into a cached template file.
func loadTemplates() {
	templates = template.Must(template.ParseFiles("templates/homepage.html", "templates/new.html"))
}

func main() {
	loadTemplates()
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
