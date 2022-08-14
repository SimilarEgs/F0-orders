package server

import (
	"net/http"
	"time"

	"github.com/SimilarEgs/L0-orders/config"
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
		WriteTimeout:   time.Second * cfg.HTTP.WriteTimeout,
		ReadTimeout:    time.Second * cfg.HTTP.ReadTimeout,
	}
	return s.httpServer.ListenAndServe()
}
