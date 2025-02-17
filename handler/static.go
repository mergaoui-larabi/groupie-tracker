package handler

import (
	"net/http"
	"os"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorTemp(w, http.StatusMethodNotAllowed, "METHOD NOT ALLOWED")
		return
	}
	filePath, err := os.Stat(r.URL.Path[1:])
	if err != nil {
		if os.IsNotExist(err) {
			ErrorTemp(w, http.StatusNotFound, "NOT FOUND")
			return
		}
		ErrorTemp(w, http.StatusInternalServerError, "INTERNAL SERVER ERROR")
		return
	}
	if filePath.IsDir() {
		ErrorTemp(w, http.StatusNotFound, "NOT FOUND")
		return
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}
