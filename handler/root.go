package handler

import (
	"html/template"
	"net/http"
	"sync"

	"grptrker/model"
)

const (
	URL1 = "https://groupietrackers.herokuapp.com/api/artists"
	URL2 = "https://groupietrackers.herokuapp.com/api/locations"
	URL3 = "https://groupietrackers.herokuapp.com/api/dates"
	URL4 = "https://groupietrackers.herokuapp.com/api/relation"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	var Artists []model.Artist
	if r.Method != http.MethodGet {
		ErrorTemp(w, http.StatusBadRequest, "METHOD NOT ALLOWED")
		return
	}
	if r.URL.Path != "/" {
		ErrorTemp(w, http.StatusNotFound, "NOT FOUND")
		return
	}
	waiter := sync.WaitGroup{}
	waiter.Add(1)
	go Fetch(URL1, &Artists, &waiter)
	waiter.Wait()
	temp := template.Must(template.ParseFiles("./static/temp/home.html"))
	err := temp.Execute(w, Artists)
	if err != nil {
		ErrorTemp(w, http.StatusInternalServerError, "INTERNAL SERVER ERROR")
		return
	}
}
