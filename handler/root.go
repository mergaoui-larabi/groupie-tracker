package handler

import (
	"grptrker/model"
	"html/template"
	"net/http"
	"strconv"
	"sync"
)

const URL1 = "https://groupietrackers.herokuapp.com/api/artists"
const URL2 = "https://groupietrackers.herokuapp.com/api/locations"
const URL3 = "https://groupietrackers.herokuapp.com/api/dates"
const URL4 = "https://groupietrackers.herokuapp.com/api/relation"

func RootHandler(w http.ResponseWriter, r *http.Request) {
	waiter := sync.WaitGroup{}
	waiter.Add(1)
	go Fetch(URL1, &model.Artists, &waiter)
	waiter.Wait()
	temp := template.Must(template.ParseFiles("./temp/home.html"))
	temp.Execute(w, model.Artists)
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	var SingleArtist model.Artist
	waiter := sync.WaitGroup{}
	id := r.PathValue("id")
	id_int, errAtoi := strconv.Atoi(id)
	if len(model.Artists) == 0 {
		waiter.Add(1)
		go Fetch(URL1, &model.Artists, &waiter)
		waiter.Wait()
	}

	if errAtoi != nil || id_int < 1 || id_int > len(model.Artists) {
		ErrorTemp(w, errAtoi)
		return
	}

	waiter.Add(4)

	go Fetch(URL1+"/"+id, &SingleArtist, &waiter)
	go Fetch(URL2+"/"+id, &SingleArtist.Location, &waiter)
	go Fetch(URL3+"/"+id, &SingleArtist.ConcertDate, &waiter)
	go Fetch(URL4+"/"+id, &SingleArtist.Relation, &waiter)

	waiter.Wait()

	temp, err4 := template.ParseFiles("./temp/artist.html")
	if err4 != nil {
		ErrorTemp(w, err4)
		return
	}
	temp.Execute(w, SingleArtist)
}
