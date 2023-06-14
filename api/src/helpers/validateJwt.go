package helpers

import (
	"github.com/golang-jwt/jwt"
	"github.com/kjblanchard/BowlingWebApp/models"
	"net/http"
	"fmt"
)

// secret key for jwt for now.
var jwtKey = []byte("my_secret_key")

func ValidateJwt(w http.ResponseWriter, cookie *http.Cookie) (*models.Claims, error) {
	// Get the JWT string from the cookie
	tknStr := cookie.Value

	// Initialize a new instance of `Claims`
	claims := &models.Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return nil, fmt.Errorf("User is unauthorized")
		}
		w.WriteHeader(http.StatusBadRequest)
		return nil, fmt.Errorf("Bad Request")
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return nil, fmt.Errorf("Token is invalid")
	}
	return claims, nil

}
