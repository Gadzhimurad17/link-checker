package middleware

import "net/http"

func ContentTypeChecker(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost || req.Method == http.MethodPatch || req.Method == http.MethodPut {
			contentType := req.Header.Get("Content-Type")
			if contentType != "application/json" {
				http.Error(rw, "Unsupported Media Type. Only application/json is allowed.", http.StatusUnsupportedMediaType)
				return
			}
		}
		next(rw,req)
	}

}
