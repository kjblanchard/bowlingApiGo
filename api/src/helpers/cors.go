package helpers

import "net/http"

// If you use * here it will allow from all sites.
var enabledWebsites = "http://bowling.supergoon.com:8080"

func EnableCors(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", enabledWebsites)
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	(*w).Header().Set("Access-Control-Allow-Headers", "X-Requested-With, X-HTTP-Method-Override, Content-Type, Accept")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
}
