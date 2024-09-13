package handlers

import (
	"net/http"
)

func SelectHandler(w http.ResponseWriter, r *http.Request) {

	isLengthStr := r.FormValue("is_length")
	var file string

	if isLengthStr == "on" {
		file = "select_length.html"
	} else {
		file = "select_weight.html"
	}

	w.Header().Set("Content-Type", "text/html")

	http.ServeFile(w, r, "./static/" + file)
	
}