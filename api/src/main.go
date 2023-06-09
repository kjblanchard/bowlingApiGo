package main

import (
	"log"
	"net/http"

	"github.com/kjblanchard/BowlingWebApp/handlers"
)

func main() {
	http.HandleFunc("/api/v1/signin", handlers.Signin)
	http.HandleFunc("/api/v1/welcome", handlers.Welcome)
	http.HandleFunc("/api/v1/refresh", handlers.Refresh)
	http.HandleFunc("/api/v1/logout", handlers.Logout)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
