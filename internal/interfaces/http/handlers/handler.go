package handlers

import (
	"encoding/json"
	"fmt"
	link "linkcheker/internal/interfaces/http/dto"
	service "linkcheker/internal/interfaces/services"
	"log"
	"net/http"
	"os"
	"strconv"
)

type LinkCheckerHandler struct {
	service *service.LinkCheckerService
}

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

func (h *LinkCheckerHandler) CheckLinks(rw http.ResponseWriter, req *http.Request) {
	linksAmount, err := strconv.Atoi(os.Getenv("MAX_LINKS_AMOUNT"))
	if err != nil || linksAmount <= 0 {
		http.Error(rw, "invalid MAX_LINKS_AMOUNT", http.StatusInternalServerError)
		return
	}

	encodeReq, err := decode[link.LinkRequest](req)
	if err != nil {
		log.Printf("Decoding error json %v", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	results := h.service.CheckLinks(
		req.Context(),
		encodeReq.Links,
	)

	rw.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(rw).Encode(results); err != nil {
		log.Printf("Encoding error: %v", err)
	}
}

func CheckApiHealth(rw http.ResponseWriter, req *http.Request) {

}
