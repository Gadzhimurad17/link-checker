package http

import (
	"encoding/json"
	"fmt"
	"linkcheker/internal/infrastructure/dto"
	"log"
	"net/http"
	"strconv"
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
	limit := req.PathValue("limit")
	limitInt, err := strconv.ParseInt(limit, 4, 8)
	if err != nil {

	}

	results := h.service.PopularLinks(limitInt)

	if err := encode(rw, http.StatusOK, results); err != nil {
		log.Printf("Encoding to json error %v", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *Handlers) CheckApiHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
