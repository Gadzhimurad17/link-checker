package handlers

import "net/http"

func (h *Handlers) CheckApiHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

