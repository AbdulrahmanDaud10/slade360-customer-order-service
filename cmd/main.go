package main

import (
	"fmt"
	"os"

	"github.com/AbdulrahmanDaud10/slade360-customer-order-service/pkg/api"
	"github.com/AbdulrahmanDaud10/slade360-customer-order-service/pkg/app"
	"github.com/AbdulrahmanDaud10/slade360-customer-order-service/pkg/repository"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config
)

func Googleinit() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

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
