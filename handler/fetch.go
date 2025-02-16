package handler

import (
	"encoding/json"
	"grptrker/model"
	"io"
	"net/http"
	"sync"
	"time"
)

func Fetch(url string, data any, wg *sync.WaitGroup) error {
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
	switch data.(type) {
	case *[]model.Artist:
		err := json.Unmarshal(body, data)
		if err != nil {
			return err
		}
	case *model.Artist:
		err := json.Unmarshal(body, data)
		if err != nil {
			return err
		}
	case *model.Location:
		err := json.Unmarshal(body, data)
		if err != nil {
			return err
		}
	case *model.Dates:
		err := json.Unmarshal(body, data)
		if err != nil {
			return err
		}
	case *model.Relation:
		err := json.Unmarshal(body, data)
		if err != nil {
			return err
		}
	default:
		return nil
	}
	return nil
}
