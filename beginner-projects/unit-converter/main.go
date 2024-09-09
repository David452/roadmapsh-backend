package main

import (
	"net/http"
	"fmt"

	"github.com/David452/unit-converter/internal/handlers"
)

const (
	PORT = ":8080"
)

func main() {

	fmt.Printf("Listening on port %s\n", PORT)

	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/", fs)

	http.HandleFunc("/convert", handlers.ConvertHandler)

	http.ListenAndServe(PORT, nil)
}