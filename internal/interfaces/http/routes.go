package http

import "net/http"

func (s *Server) RegisterRoutes(mux *http.ServeMux) {
	m := s.middleware
	mux.HandleFunc("POST /api/links", m.ValidateApiKey(m.ContentTypeChecker(s.handlers.CheckLinks)))
	//s.GetMux().HandleFunc("GET/api/health", hdlr.CheckApiHealth)
}
