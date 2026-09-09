package server

import "net/http"

type Server struct {
	mux    *http.ServeMux
	server http.Server
}

func (s *Server) GetMux() *http.ServeMux {
	return s.mux
}
func StartServer(s *Server) {

}
