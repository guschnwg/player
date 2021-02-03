package server

import "net/http"

func enableCors(w *http.ResponseWriter, r *http.Request) {
	if r.Referer() == "https://guschnwg.github.io/" {
		(*w).Header().Set("Access-Control-Allow-Origin", "https://guschnwg.github.io/*")
	} else if r.Referer() == "https://evening-ridge-00695.herokuapp.com/*" {
		(*w).Header().Set("Access-Control-Allow-Origin", "https://evening-ridge-00695.herokuapp.com/*")
	}
}
