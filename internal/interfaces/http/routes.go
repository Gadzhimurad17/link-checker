package http

import "net/http"

func (s *Server) RegisterRoutes(mux *http.ServeMux) {
	m := s.middleware
	mux.HandleFunc("POST /api/links", m.ValidateApiKey(m.ContentTypeChecker(s.handlers.CheckLinks)))
	mux.HandleFunc("GET /api/links/{limit}/{offset}", s.handlers.GetMostPopularLinks)

	mux.HandleFunc("GET /api/health", s.handlers.CheckApiHealth)
}
