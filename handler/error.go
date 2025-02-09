package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func ErrorTemp(w http.ResponseWriter, err error) {
	temp, errParse := template.ParseFiles("./temp/error.html")
	if errParse != nil {
		fmt.Println(errParse)
		return
	}
	temp.Execute(w, err)

}
