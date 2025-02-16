package handler

import (
	"net/http"
	"os"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorTemp(w, http.StatusMethodNotAllowed)
		return
	}
	filePath, err := os.Stat(r.URL.Path[1:])
	if err != nil {
		if os.IsNotExist(err) {
			ErrorTemp(w, http.StatusNotFound)
			return
		}
		ErrorTemp(w, http.StatusInternalServerError)
		return
	}
	if filePath.IsDir() {
		ErrorTemp(w, http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}
