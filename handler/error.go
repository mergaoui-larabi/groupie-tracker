package handler

import (
	"html/template"
	"net/http"
)

type errorcontent struct {
	StatusCode int
	Message    string
}

func ErrorTemp(w http.ResponseWriter, code int, errormsg string) {
	temp, err := template.ParseFiles("./static/temp/error.html")

	if err != nil {
		http.Error(w, "internal server error : unable to render template", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)

	data := errorcontent{
		StatusCode: code,
		Message:    errormsg,
	}
	err = temp.Execute(w, data)
	if err != nil {
		http.Error(w, "internal server error : unable to render template", http.StatusInternalServerError)
		return
	}
}
