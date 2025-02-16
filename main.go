package main

import (
	"fmt"
	"grptrker/handler"
	"net/http"
)

func main() {

	// router := http.NewServeMux()
	http.HandleFunc("/", handler.RootHandler)
	http.HandleFunc("/artist/{id}", handler.ArtistHandler)
	http.HandleFunc("/static/", handler.StaticHandler)
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
