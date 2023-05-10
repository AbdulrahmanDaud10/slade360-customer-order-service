package app

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
}

func NewServer(router *mux.Router) *server {
	return &server{
		router: router,
	}
}

func (s *server) Run() error {

	LISTEN_ADDRESS := os.Getenv("LISTEN_ADDRESS")
	LISTEN_PORT := os.Getenv("LISTEN_PORT")
	ListenADDR := LISTEN_ADDRESS + ":" + LISTEN_PORT

	serve := &http.Server{
		Addr:              ListenADDR,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}
	log.Print("App Running")
	serve.ListenAndServe()

	return nil
}
