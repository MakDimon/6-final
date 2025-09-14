package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Strct struct {
	Logr *log.Logger
	Hts  http.Server
}

// Srvr - функция сервера
func Srvr(logger *log.Logger) Strct {
	r := http.NewServeMux()
	r.HandleFunc("/", handlers.HandleRoot)
	r.HandleFunc("/upload", handlers.HandleUpload)

	return Strct{
		Logr: logger,
		Hts: http.Server{
			Addr:         ":8080",
			Handler:      r,
			ErrorLog:     logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
}

// Serve - запускает сервер
func (s *Strct) Serve() {
	if err := s.Hts.ListenAndServe(); err != nil {
		s.Logr.Fatal(err)
	}
}
