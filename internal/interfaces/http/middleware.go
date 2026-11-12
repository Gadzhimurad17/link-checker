package http

import (
	"context"
	"net/http"
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

func (m *Middleware) WithGracefulShutdown(next http.Handler, cancelCh <-chan struct{}) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx, cancel := m.cancel(req.Context(), cancelCh)
		defer cancel()

		req = req.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}

func (m *Middleware) cancel(parentCtx context.Context, cancelCh <-chan struct{}) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(parentCtx)

	go func() {
		select {
		case <-cancelCh:
			cancel()
		case <-parentCtx.Done():
		}
	}()

	return ctx, cancel
}
