package server

import (
	"net/http"

	"github.com/SimilarEgs/F0-orders/config"
)

const (
	maxHeaderBytes = 1 << 20
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) RunServer(handler http.Handler, cfg *config.Config) error {

	s.httpServer = &http.Server{
		Addr:           cfg.HTTP.Port,
		Handler:        handler,
		MaxHeaderBytes: maxHeaderBytes,
		WriteTimeout:   cfg.HTTP.WriteTimeout,
		ReadTimeout:    cfg.HTTP.ReadTimeout,
	}
	return s.httpServer.ListenAndServe()
}
