package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
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
	// start server
	if err := server.Run(); err != nil {
		return err
	}
	return nil
}
