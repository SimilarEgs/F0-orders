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

	// serving html
	fs := http.FileServer(http.Dir("./static/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	http.HandleFunc("/orders", OrderByIdHandler)

	// serving favicon
	is := http.FileServer(http.Dir("./static/templates"))
	http.Handle("/templates/", http.StripPrefix("/templates", is))
	http.HandleFunc("/favicon.ico", faviconHandler)

	

	s.httpServer = &http.Server{
		Addr:           cfg.HTTP.Port,
		MaxHeaderBytes: maxHeaderBytes,
		WriteTimeout:   time.Second * cfg.HTTP.WriteTimeout,
		ReadTimeout:    time.Second * cfg.HTTP.ReadTimeout,
	}
	return s.httpServer.ListenAndServe()
}
