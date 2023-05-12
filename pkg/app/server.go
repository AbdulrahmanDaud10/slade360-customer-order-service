package app

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AbdulrahmanDaud10/slade360-customer-order-service/pkg/api"
	"github.com/gorilla/mux"
)

type server struct {
	router          *mux.Router
	customerservice api.CustomerService
}

func NewServer(router *mux.Router, customerservice api.CustomerService) *server {
	return &server{
		router:          router,
		customerservice: customerservice,
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
