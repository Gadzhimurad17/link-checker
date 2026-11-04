package http

import (
	"encoding/json"
	"fmt"
	"linkcheker/internal/infrastructure/dto"
	"log"
	"net/http"
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

func (h *Handlers) CheckLinks(rw http.ResponseWriter, req *http.Request) {

	decodeReq, err := decode[dto.LinkRequest](req)
	if err != nil {
		log.Printf("Decoding error json %v", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	results := h.service.CheckLinks(
		req.Context(),
		decodeReq.Links,
	)

	if err := encode(rw, http.StatusOK, results); err != nil {
		log.Printf("Encoding to json error %v", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

}

func (h *Handlers) GetMostPopularLinks(rw http.ResponseWriter, req *http.Request) {
	results := h.service.PopularLinks()
	if err := encode(rw, http.StatusOK, results); err != nil {
		log.Printf("Encoding to json error %v", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}
