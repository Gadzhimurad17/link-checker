package http

import "net/http"

type middleware struct {
	apiKey string
}

func (m *middleware) ValidateApiKey(next http.HandlerFunc) http.HandlerFunc {

	return func(rw http.ResponseWriter, req *http.Request) {
		apiKey := req.Header.Get("x-api-key")
		if apiKey != m.apiKey {
			http.Error(rw, "Forbidden", http.StatusForbidden)
			return
		}
		next(rw, req)
	}
}

func (m *middleware) ContentTypeChecker(next http.HandlerFunc) http.HandlerFunc {
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
