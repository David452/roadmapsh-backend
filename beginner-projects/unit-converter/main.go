package main

import (
	"net/http"
	"fmt"
)

const (
	PORT = ":8080"
)

func main() {

	fmt.Printf("Listening on port %s\n", PORT)

	fs := http.FileServer(http.Dir("./web/static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/templates/index.html")
	})

	http.ListenAndServe(PORT, nil)
}