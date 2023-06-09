package main

import (
	"log"
	"net/http"

	"github.com/kjblanchard/BowlingWebApp/handlers"
	"github.com/kjblanchard/BowlingWebApp/helpers"
)

func main() {
	helpers.Connect()
	http.HandleFunc("/api/v1/signin", handlers.Signin)
	http.HandleFunc("/api/v1/welcome", handlers.Welcome)
	http.HandleFunc("/api/v1/refresh", handlers.Refresh)
	http.HandleFunc("/api/v1/logout", handlers.Logout)
	http.HandleFunc("/api/v1/scores", handlers.AddScore)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
