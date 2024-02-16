package server

import (
	se "AuthService/pkg/error"
	"encoding/json"
	"errors"
	"net/http"
)

func (s *HttpServer) GetTokens(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	s.logger.Info.Println("Invoked GetTokens on server. id = ", id)

	tokens, err := s.services.AuthService.GenerateTokens(id)

	if err != nil {
		s.logger.Err.Println(err.Error())
		var rErr *se.ResponseError
		if errors.As(err, &rErr) {
			se.HttpResponseError(w, *rErr)
		} else {
			se.HttpResponseError(w, se.ResponseError{Message: err.Error(), Code: http.StatusInternalServerError})
		}

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(tokens)
}

func (s *HttpServer) RefreshTokens(w http.ResponseWriter, r *http.Request) {

	refresh := r.URL.Query().Get("token")
	id := r.Header.Get("id")

	s.logger.Info.Println("Invoked RefreshTokens on server. token = ", refresh)

	tokens, err := s.services.AuthService.RefreshTokens(refresh, id)

	if err != nil {
		s.logger.Err.Println(err.Error())
		var rErr *se.ResponseError
		if errors.As(err, &rErr) {
			se.HttpResponseError(w, *rErr)
		} else {
			se.HttpResponseError(w, se.ResponseError{Message: err.Error(), Code: http.StatusInternalServerError})
		}

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(tokens)
}
