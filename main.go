package main

import (
	"fmt"
	"log"
	"net/http"

	"grptrker/handler"
)

func main() {
	// router := http.NewServeMux()
	http.HandleFunc("/", handler.RootHandler)
	http.HandleFunc("/artist/{id}", handler.ArtistHandler)
	http.HandleFunc("/static/", handler.StaticHandler)
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
