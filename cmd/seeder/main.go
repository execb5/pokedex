package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/execb5/pokedex/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func dbConfig() *gorm.DB {
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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	return db
}

func readCsvFile(filepath string) [][]string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Unable to read input file "+filepath, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as csv for "+filepath, err)
	}

	// Drop header
	return records[1:]
}

func parseType(values []string) models.Type {
	id, err := strconv.Atoi(values[0])
	if err != nil {
		log.Fatal("Unable to parse value as integer for "+values[0], err)
	}
	generationId, err := strconv.Atoi(values[2])
	if err != nil {
		log.Fatal("Unable to parse value as integer for "+values[0], err)
	}
	damageClassId, err := strconv.Atoi(values[2])
	if err != nil {
		log.Fatal("Unable to parse value as integer for "+values[0], err)
	}

	t := models.Type{
		ID:            uint(id),
		Identifier:    values[1],
		GenerationId:  int64(generationId),
		DamageClassId: int64(damageClassId),
	}
	return t
}

func main() {
	db := dbConfig()
	var records [][]string

	records = readCsvFile("data/types.csv")
	for _, record := range records {
		myType := parseType(record)
		db.Save(&myType)
	}

	// records = readCsvFile("data/pokemon_types.csv.csv")
	// for _, record := range records {
	// 	myType := parseType(record)
	// 	db.Save(&myType)
	// }
}
