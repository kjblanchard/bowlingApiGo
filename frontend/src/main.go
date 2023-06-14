package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// Template data we will pass into the templates
type TemplateData struct {
	Title       string
	ScriptFiles []string
	CssFiles []string
}

// The cached templates
var templates *template.Template

// Main page handler.
func handler(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{
		Title: "Bowling",
		ScriptFiles: []string{
			"login.js",
			"homepage.js",
			"addScore.js",
		},
		CssFiles: []string{},
	}
	err := templates.ExecuteTemplate(w, "homepage.html", data)
	if err != nil {
		log.Printf("Something happened:\n %s", err.Error())
		fmt.Fprintf(w, "Failed to load template content properly.")
	}
}

func loadTemplates() {
	// Load all the templates from disk into a pointer, and panics if it cannot be loaded.
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	loadTemplates()
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
