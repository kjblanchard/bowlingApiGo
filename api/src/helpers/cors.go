package helpers

import "net/http"

// If you use * here it will allow from all sites.
var enabledWebsites = "http://localhost:8080"

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", enabledWebsites)
}
