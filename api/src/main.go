package main

import (
	"log"
	"net/http"

	"github.com/kjblanchard/BowlingWebApp/handlers"
)

func main() {
	// we will implement these handlers in the next sections
	http.HandleFunc("/signin", handlers.Signin)
	http.HandleFunc("/welcome", handlers.Welcome)
	http.HandleFunc("/refresh", handlers.Refresh)
	http.HandleFunc("/logout", handlers.Logout)

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// var jwtKey = []byte("my_secret_key")
