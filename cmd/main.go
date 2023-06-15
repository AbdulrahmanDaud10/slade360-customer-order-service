package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AbdulrahmanDaud10/slade360-customer-order-service/pkg/api"
	"github.com/AbdulrahmanDaud10/slade360-customer-order-service/pkg/app"
	"github.com/AbdulrahmanDaud10/slade360-customer-order-service/pkg/repository"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func Run() error {
	db, err := repository.SetUpDatabaseConnection()
	if err != nil {
		return err
	}

	storage := repository.NewStorage(db)
	customerService := api.NewCustomerService(storage)

	server := app.NewServer(mux.NewRouter(), customerService)

	// start server
	if err := server.Run(); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	if err := Run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is a startup error: %s\n", err)
		os.Exit(1)
	}

	auth, err := api.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	rtr := app.New(auth)

	log.Print("Server listening on http://localhost:3000/")
	if err := http.ListenAndServe("127.0.0.1:3000", rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
