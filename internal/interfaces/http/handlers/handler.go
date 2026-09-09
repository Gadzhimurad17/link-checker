package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	//"github.com/joho/godotenv"
	link "linkcheker/internal/interfaces/http/dto"
)

func encode[T any](rw http.ResponseWriter, status int, v T) error {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	if err := json.NewEncoder(rw).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func decode[T any](req *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(req.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}
func CheckLinks(rw http.ResponseWriter, req *http.Request) {
	resultch := make(chan link.LinkResponse)
	sem := make(chan struct{}, 20)
	encodeReq, err := decode[link.LinkRequest](req)
	if err != nil {
		log.Printf("Decoding error json", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	linksAmmount := len(encodeReq.Links)
	var wg sync.WaitGroup
	for i := 0; i < linksAmmount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			status := CheckLink(req.Context(), encodeReq.Links[i])
			resultch<-link.LinkResponse{
				URL:encodeReq.Links[i],
				Status: status,
			}
		}(i)
	}
	wg.Wait()
}

func CheckLink(parentCtx context.Context, url string) bool {
	context, cancel := context.WithTimeout(parentCtx, 2*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(context, http.MethodGet, "https://"+url, nil)
	if err != nil {
		return false
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode >= 200 && resp.StatusCode < 400
}

func HandleRequest(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-type", "application/json")
}

func CheckApiHealth(rw http.ResponseWriter, req *http.Request) {

}
