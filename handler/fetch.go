package handler

import (
	"encoding/json"
	"grptrker/model"
	"io"
	"net/http"
	"sync"
)

func Fetch(url string, data any, wg *sync.WaitGroup) error {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err2 := io.ReadAll(resp.Body)
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
