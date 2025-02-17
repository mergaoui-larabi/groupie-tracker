package handler

import (
	"grptrker/model"
	"html/template"
	"net/http"
	"strconv"
	"sync"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorTemp(w, http.StatusBadRequest, "METHOD NOT ALLOWED")
		return
	}
	var SingleArtist model.Artist
	waiter := sync.WaitGroup{}
	id := r.PathValue("id")
	if id == "" {
		ErrorTemp(w, http.StatusBadRequest, "BAD REQUEST")
		return
	}
	id_int, errAtoi := strconv.Atoi(id)
	if len(model.Artists) == 0 {
		waiter.Add(1)
		go Fetch(URL1, &model.Artists, &waiter)
		waiter.Wait()
	}

	if errAtoi != nil || id_int < 1 || id_int > len(model.Artists) {
		ErrorTemp(w, http.StatusNotFound, "NOT FOUND")
		return
	}

	waiter.Add(4)

	go Fetch(URL1+"/"+id, &SingleArtist, &waiter)
	go Fetch(URL2+"/"+id, &SingleArtist.Location, &waiter)
	go Fetch(URL3+"/"+id, &SingleArtist.ConcertDate, &waiter)
	go Fetch(URL4+"/"+id, &SingleArtist.Relation, &waiter)

	waiter.Wait()

	temp, err4 := template.ParseFiles("./static/temp/artist.html")
	if err4 != nil {
		ErrorTemp(w, http.StatusInternalServerError, "INTERNAL SERVER ERROR")
		return
	}
	err := temp.Execute(w, SingleArtist)
	if err != nil {
		ErrorTemp(w, http.StatusInternalServerError, "INTERNAL SERVER ERROR")
		return
	}
}
