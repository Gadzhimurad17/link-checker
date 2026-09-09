package middleware

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
