package router

import (
	hdlr "linkcheker/internal/interfaces/http/handlers"
	srv "linkcheker/internal/interfaces/http/server"
)

func Routes(s *srv.Server) {
	s.GetMux().HandleFunc("POST/links", hdlr.CheckLinks)
	s.GetMux().HandleFunc("GET/links")
	s.GetMux().HandleFunc("GET/api/health", hdlr.CheckApiHealth)
}
