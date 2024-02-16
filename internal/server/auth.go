package server

import (
	server_error "AuthService/pkg/error"
	"encoding/json"
	"net/http"
)

func (s *HttpServer) GetTokens(w http.ResponseWriter, r *http.Request) {

	s.logger.Info.Println("Invoked GetTokens on server")

	id := r.URL.Query().Get("id")

	tokens, err := s.services.AuthService.GenerateTokens(id)

	if err != nil {
		s.logger.Err.Println(err.Error())
		server_error.HttpResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(tokens)
}

func (s *HttpServer) RefreshTokens(w http.ResponseWriter, r *http.Request) {
	s.logger.Info.Println("Invoked RefreshTokens on server")

	refresh := r.URL.Query().Get("token")
	id := r.Header.Get("id")

	tokens, err := s.services.AuthService.RefreshTokens(refresh, id)
	if err != nil {
		s.logger.Err.Println(err.Error())
		server_error.HttpResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(tokens)
}
