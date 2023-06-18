package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/kjblanchard/BowlingWebApp/helpers"
	"github.com/kjblanchard/BowlingWebApp/models"
	"log"
	"net/http"
)


func AddScore(w http.ResponseWriter, r *http.Request) {
	helpers.EnableCors(&w, r)
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	claims, validationError := helpers.ValidateJwt(w, c)
	if validationError != nil {
		log.Printf("Bad validation for JWT")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userId := claims.UserId
	var score models.Score
	jsonError := json.NewDecoder(r.Body).Decode(&score)
	if jsonError != nil {

		fmt.Printf("Failed parsing body: %s, \nError: %v", r.Body, jsonError)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("No")
	game, error := helpers.AddScore(userId, score.Score)
	if error != nil {
		log.Printf("Failure with error %s", error)
		fmt.Printf("Failed to add score: %v", error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("Wut")
	e := json.NewEncoder(w).Encode(&game)
	if e != nil {
		fmt.Printf("Failed to encode %s", e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
