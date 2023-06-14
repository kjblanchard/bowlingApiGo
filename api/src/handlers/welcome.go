package handlers

import (
	"fmt"
	"log"
	"net/http"
	"github.com/kjblanchard/BowlingWebApp/helpers"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	// We can obtain the session token from the requests cookies, which come with every request
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
		return
	}

	// Finally, return the welcome message to the user, along with their
	// username given in the token
	w.Write([]byte(fmt.Sprintf("Welcome %s %d!", claims.Username, claims.UserId)))
}
