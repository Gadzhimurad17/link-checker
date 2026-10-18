package http

import "net/http"

func (s *Server) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/links", s.handlers.CheckLinks)
	//s.GetMux().HandleFunc("GET/api/health", hdlr.CheckApiHealth)
}
