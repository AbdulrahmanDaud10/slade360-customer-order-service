package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *server) Routes() *mux.Router {
	router := s.router

	http.HandleFunc("/", handleMain)
	http.HandleFunc("/login", handleGoogleLogin)
	http.HandleFunc("/callback", handleGoogleCallback)

	return router
}
