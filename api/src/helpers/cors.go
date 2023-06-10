package helpers

import "net/http"

// var enabledWebsites = "*"
var enabledWebsites = "http://localhost:8080"

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", enabledWebsites)
}
