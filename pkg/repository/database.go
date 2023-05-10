package repository

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	err error
)

// // SetupDB opens a database and saves the reference to `Database` struct.
func SetUpDatabaseConnection() (*gorm.DB, error) {
	var db *gorm.DB

	driver := os.Getenv("DB_DRIVER")
	database := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("driver=%s database=%s username=%s host=%s password=%s port=%s sslmode=disable",
		driver,
		database,
		username,
		host,
		password,
		port,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: ", err)
		return nil, err
	}
	// TODO: run the customer migration after implemented
	// db.AutoMigrate(api.customer{})
	// if err != nil {
	// 	log.Fatal("Error connecting to the database...", err)
	// }

	fmt.Println("Database connection successful...")

	return db, nil
}
