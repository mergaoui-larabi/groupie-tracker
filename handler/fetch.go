package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"time"

	"grptrker/model"
)

type atawix interface {
	*[]model.Artist | *model.Location | *model.Artist | *model.Dates | *model.Relation
}

func Fetch[T atawix](url string, data T, wg *sync.WaitGroup) error {
	Client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := Client.Get(url)
	if err != nil {
		return err
	}
	body, err2 := io.ReadAll(resp.Body)
	defer func() {
		wg.Done()
		resp.Body.Close()
	}()
	if err2 != nil {
		return err2
	}
	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}
	return nil
}
