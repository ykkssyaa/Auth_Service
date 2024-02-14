package server

import (
	"AuthService/internal/service"
	"AuthService/pkg/logger"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type HttpServer struct {
	services *service.Services
	logger   *logger.Logger
}

func NewHttpServer(services *service.Services, logger *logger.Logger, addr string) *http.Server {
	server := &HttpServer{services: services, logger: logger}

	r := mux.NewRouter()
	r.HandleFunc("/auth", server.GetTokens).Methods("GET")
	r.HandleFunc("/refresh", server.RefreshTokens).Methods("GET")

	return &http.Server{
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Addr:           addr,
	}
}
