package server

import "net/http"

func (s *HttpServer) GetTokens(w http.ResponseWriter, r *http.Request) {

	s.logger.Info.Println("Invoked GetTokens on server")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}

func (s *HttpServer) RefreshTokens(w http.ResponseWriter, r *http.Request) {
	s.logger.Info.Println("Invoked RefreshTokens on server")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}
