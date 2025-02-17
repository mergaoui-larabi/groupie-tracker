package handler

import (
	"grptrker/model"
	"html/template"
	"net/http"
	"sync"
)

const URL1 = "https://groupietrackers.herokuapp.com/api/artists"
const URL2 = "https://groupietrackers.herokuapp.com/api/locations"
const URL3 = "https://groupietrackers.herokuapp.com/api/dates"
const URL4 = "https://groupietrackers.herokuapp.com/api/relation"

func RootHandler(w http.ResponseWriter, r *http.Request) {
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
	go Fetch(URL1, &model.Artists, &waiter)
	waiter.Wait()
	temp := template.Must(template.ParseFiles("./static/temp/home.html"))
	err := temp.Execute(w, model.Artists)
	if err != nil {
		ErrorTemp(w, http.StatusInternalServerError, "INTERNAL SERVER ERROR")
		return
	}
}
