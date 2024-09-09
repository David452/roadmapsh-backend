package handlers

import (
	"net/http"
)

func CSSHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/css/style.css")
}