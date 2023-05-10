package main

import (
	"fmt"
	"os"

	"github.com/AbdulrahmanDaud10/slade360-customer-order-service/pkg/app"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		panic(err)
	}
}

func main() {
	if err := Run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is a startup error: %s\n", err)
		os.Exit(1)
	}
}
func Run() error {
	server := app.NewServer(mux.NewRouter())

	// start server
	if err := server.Run(); err != nil {
		return err
	}
	return nil
}
