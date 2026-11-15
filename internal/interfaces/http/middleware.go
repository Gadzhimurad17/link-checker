package http

import (
	"log"
	"net/http"
	"time"

	"github.com/Gadzhimurad17/link-checker/internal/interfaces/http/utils"
)

type Middleware struct {
	apiKey string
}

func NewMiddleware(apiKey string) *Middleware {
	return &Middleware{
		apiKey: apiKey,
	}
}

func (m *Middleware) ValidateApiKey(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		apiKey := req.Header.Get("X-Api-Key")
		if apiKey != m.apiKey {
			http.Error(rw, "Forbidden", http.StatusForbidden)
			return
		}
		next(rw, req)
	}
}

func (m *Middleware) ContentTypeChecker(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost || req.Method == http.MethodPatch || req.Method == http.MethodPut {
			contentType := req.Header.Get("Content-Type")
			if contentType != "application/json" {
				http.Error(rw, "Unsupported Media Type. Only application/json is allowed.", http.StatusUnsupportedMediaType)
				return
			}
		}
		next(rw, req)
	}
}

func (m *Middleware) LogRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		recorder := &utils.ResponseRecorder{ResponseWriter: w}

		next.ServeHTTP(recorder, r)

		status := recorder.GetStatus()
		if status == 0 {
			status = http.StatusOK
		}

		color := "\033[32m"
		switch {
		case status >= http.StatusInternalServerError:
			color = "\033[31m"
		case status >= http.StatusBadRequest:
			color = "\033[33m"
		case status >= http.StatusMultipleChoices:
			color = "\033[36m"
		}

		log.Printf(
			"%s%d %-7s\033[0m | %10s | %s",
			color,
			status,
			r.Method,
			time.Since(start).Round(time.Microsecond),
			r.URL.Path,
		)
	})
}
