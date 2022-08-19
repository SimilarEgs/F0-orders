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

func (s *Server) RunServer(cfg *config.Config) error {

	// serving static files
	fs := http.FileServer(http.Dir("./static/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	http.HandleFunc("/orders", OrderByIdHandler)

	s.httpServer = &http.Server{
		Addr:           cfg.HTTP.Port,
		MaxHeaderBytes: maxHeaderBytes,
		WriteTimeout:   time.Second * cfg.HTTP.WriteTimeout,
		ReadTimeout:    time.Second * cfg.HTTP.ReadTimeout,
	}
	return s.httpServer.ListenAndServe()
}
