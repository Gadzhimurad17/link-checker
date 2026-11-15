package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Gadzhimurad17/link-checker/internal/domain/rules"
	"github.com/Gadzhimurad17/link-checker/internal/infrastructure/dto"
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
	rules.ValidateLinks(decodeReq)

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
		http.Error(rw, "Invalid limit", http.StatusBadRequest)
		return
	}

	offset := req.PathValue("offset")
	offsetInt, err := strconv.ParseInt(offset, 4, 8)
	if err != nil {
		http.Error(rw, "Invalid limit", http.StatusBadRequest)
		return
	}

	results, err := h.service.PopularLinks(limitInt, offsetInt)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := encode(rw, http.StatusOK, results); err != nil {
		log.Printf("Encoding to json error %v", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}
