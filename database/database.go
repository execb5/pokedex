package database

import (
	"fmt"
	"os"

	"github.com/execb5/pokedex/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Initialize() {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "pokedex"
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "pokedex"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "pokedex"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	dbSslMode := os.Getenv("DB_SSL_MODE")
	if dbSslMode == "" {
		dbSslMode = "disable"
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", dbHost, dbUser, dbPassword, dbName, dbPort, dbSslMode)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db = database
}

func InitializeTestDB() {
	if db != nil {
		return
	}
	database, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database")
	}

	database.AutoMigrate(&models.Pokemon{})

	db = database
}

func GetDB() *gorm.DB {
	return db
}
