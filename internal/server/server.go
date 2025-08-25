package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type ServerMy struct {
	Loger  *log.Logger
	Server http.Server
}

func LaunchServer(logger *log.Logger) *ServerMy {
	router := http.NewServeMux()

	router.HandleFunc("/", handlers.HtmlHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	serv := &ServerMy{
		Loger: logger,
		Server: http.Server{
			Addr:         ":8080",
			Handler:      router,
			ErrorLog:     logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
	return serv
}

func StartServer(s *ServerMy) error {
	err := http.ListenAndServe(s.Server.Addr, s.Server.Handler)
	return err
}
