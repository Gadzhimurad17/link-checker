package utils

import (
	"net/http"
)

type ResponseRecorder struct {
	http.ResponseWriter

	status int
}

func (w *ResponseRecorder) GetStatus() int {
	return w.status
}
func (w *ResponseRecorder) WriteHeader(status int) {
	if status >= 100 && status < 200 {
		w.ResponseWriter.WriteHeader(status)
		return
	}

	if w.status != 0 {
		return
	}

	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *ResponseRecorder) Write(data []byte) (int, error) {
	if w.status == 0 {
		w.WriteHeader(http.StatusOK)
	}

	return w.ResponseWriter.Write(data)
}

func (w *ResponseRecorder) Unwrap() http.ResponseWriter {
	return w.ResponseWriter
}
