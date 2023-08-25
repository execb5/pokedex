package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/execb5/pokedex/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
		log.Fatal("[type][id] Unable to parse value as integer for "+values[0], err)
	}
	generationId, err := strconv.Atoi(values[2])
	if err != nil {
		log.Fatal("[type][generationId] Unable to parse value as integer for "+values[0], err)
	}
	damageClassId, err := strconv.Atoi(values[2])
	if err != nil {
		log.Fatal("[type][damageClassId] Unable to parse value as integer for "+values[0], err)
	}

	t := models.Type{
		ID:            uint(id),
		Identifier:    values[1],
		GenerationId:  int64(generationId),
		DamageClassId: int64(damageClassId),
	}
	return t
}

func parseShape(values []string) models.Shape {
	id, err := strconv.Atoi(values[0])
	if err != nil {
		log.Fatal("[shape][id] Unable to parse value as integer for "+values[0], err)
	}

	t := models.Shape{
		ID:         uint(id),
		Identifier: values[1],
	}
	return t
}

func parseColor(values []string) models.Color {
	id, err := strconv.Atoi(values[0])
	if err != nil {
		log.Fatal("[color][id] Unable to parse value as integer for "+values[0], err)
	}

	t := models.Color{
		ID:         uint(id),
		Identifier: values[1],
	}
	return t
}

func parseSpecies(values []string) models.Species {
	id, err := strconv.Atoi(values[0])
	if err != nil {
		log.Fatal("[species][id] Unable to parse value as integer for "+values[0], err)
	}
	generationId, err := strconv.Atoi(values[2])
	if err != nil {
		log.Fatal("[species][generationId] Unable to parse value as integer for "+values[2], err)
	}
	var evolvesFromSpeciesId sql.NullInt64
	if values[3] == "" {
		evolvesFromSpeciesId = sql.NullInt64{
			Int64: 0,
			Valid: false,
		}
	} else {
		value, err := strconv.Atoi(values[3])
		if err != nil {
			log.Fatal("[species][evolvesFromSpeciesId] Unable to parse value as integer for "+values[3], err)
		}
		evolvesFromSpeciesId = sql.NullInt64{
			Int64: int64(value),
			Valid: true,
		}
	}
	evolutionChainId, err := strconv.Atoi(values[4])
	if err != nil {
		log.Fatal("[species][evolutionChainId] Unable to parse value as integer for "+values[4], err)
	}
	colorId, err := strconv.Atoi(values[5])
	if err != nil {
		log.Fatal("[species][colorId] Unable to parse value as integer for "+values[5], err)
	}
	var shapeId sql.NullInt64
	if values[6] == "" {
		shapeId = sql.NullInt64{
			Int64: 0,
			Valid: false,
		}
	} else {
		value, err := strconv.Atoi(values[6])
		if err != nil {
			log.Fatal("[species][shapeId] Unable to parse value as integer for "+values[6], err)
		}
		shapeId = sql.NullInt64{
			Int64: int64(value),
			Valid: true,
		}
	}
	var habitatId sql.NullInt64
	if values[7] == "" {
		habitatId = sql.NullInt64{
			Int64: 0,
			Valid: false,
		}
	} else {
		value, err := strconv.Atoi(values[7])
		if err != nil {
			log.Fatal("[species][habitatId] Unable to parse value as integer for "+values[7], err)
		}
		habitatId = sql.NullInt64{
			Int64: int64(value),
			Valid: true,
		}
	}
	genderRate, err := strconv.Atoi(values[8])
	if err != nil {
		log.Fatal("[species][genderRate] Unable to parse value as integer for "+values[8], err)
	}
	captureRate, err := strconv.Atoi(values[9])
	if err != nil {
		log.Fatal("[species][captureRate] Unable to parse value as integer for "+values[9], err)
	}
	baseHappiness, err := strconv.Atoi(values[10])
	if err != nil {
		log.Fatal("[species][baseHappiness] Unable to parse value as integer for "+values[10], err)
	}
	isBaby, err := strconv.ParseBool(values[11])
	if err != nil {
		log.Fatal("[species][isBaby] Unable to parse value as bool for "+values[11], err)
	}
	hatchCounter, err := strconv.Atoi(values[12])
	if err != nil {
		log.Fatal("[species][hatchCounter] Unable to parse value as integer for "+values[12], err)
	}
	hasGenderDifferences, err := strconv.ParseBool(values[13])
	if err != nil {
		log.Fatal("[species][hasGenderDifferences] Unable to parse value as bool for "+values[13], err)
	}
	growthRateId, err := strconv.Atoi(values[14])
	if err != nil {
		log.Fatal("[species][growthRateId] Unable to parse value as integer for "+values[14], err)
	}
	formsSwitchable, err := strconv.ParseBool(values[15])
	if err != nil {
		log.Fatal("[species][formsSwitchable] Unable to parse value as bool for "+values[15], err)
	}
	isLegendary, err := strconv.ParseBool(values[16])
	if err != nil {
		log.Fatal("[species][isLegendary] Unable to parse value as bool for "+values[16], err)
	}
	isMythical, err := strconv.ParseBool(values[17])
	if err != nil {
		log.Fatal("[species][isMythical] Unable to parse value as bool for "+values[17], err)
	}
	order, err := strconv.Atoi(values[18])
	if err != nil {
		log.Fatal("[species][order] Unable to parse value as integer for "+values[18], err)
	}
	var conquestOrder sql.NullInt64
	if values[19] == "" {
		conquestOrder = sql.NullInt64{
			Int64: 0,
			Valid: false,
		}
	} else {
		value, err := strconv.Atoi(values[19])
		if err != nil {
			log.Fatal("[species][conquestOrder] Unable to parse value as integer for "+values[19], err)
		}
		conquestOrder = sql.NullInt64{
			Int64: int64(value),
			Valid: true,
		}
	}

	t := models.Species{
		ID:                   uint(id),
		Identifier:           values[1],
		GenerationId:         generationId,
		EvolvesFromSpeciesId: evolvesFromSpeciesId,
		EvolutionChainId:     evolutionChainId,
		ColorId:              uint(colorId),
		ShapeId:              shapeId,
		HabitatId:            habitatId,
		GenderRate:           genderRate,
		CaptureRate:          captureRate,
		BaseHappiness:        baseHappiness,
		IsBaby:               isBaby,
		HatchCounter:         hatchCounter,
		HasGenderDifferences: hasGenderDifferences,
		GrowthRateId:         growthRateId,
		FormsSwitchable:      formsSwitchable,
		IsLegendary:          isLegendary,
		IsMythical:           isMythical,
		Order:                order,
		ConquestOrder:        conquestOrder,
	}
	return t
}

func parsePokemon(values []string) models.Pokemon {
	id, err := strconv.Atoi(values[0])
	if err != nil {
		log.Fatal("[pokemon][id] Unable to parse value as integer for "+values[0], err)
	}
	speciesId, err := strconv.Atoi(values[2])
	if err != nil {
		log.Fatal("[pokemon][speciesId] Unable to parse value as integer for "+values[2], err)
	}
	height, err := strconv.Atoi(values[3])
	if err != nil {
		log.Fatal("[pokemon][height] Unable to parse value as integer for "+values[3], err)
	}
	weight, err := strconv.Atoi(values[4])
	if err != nil {
		log.Fatal("[pokemon][weight] Unable to parse value as integer for "+values[4], err)
	}
	baseExperience, err := strconv.Atoi(values[5])
	if err != nil {
		log.Fatal("[pokemon][baseExperience] Unable to parse value as integer for "+values[5], err)
	}
	order, err := strconv.Atoi(values[6])
	if err != nil {
		log.Fatal("[pokemon][order] Unable to parse value as integer for "+values[6], err)
	}
	isDefault, err := strconv.ParseBool(values[7])
	if err != nil {
		log.Fatal("[pokemon][isDefault] Unable to parse value as integer for "+values[7], err)
	}

	t := models.Pokemon{
		ID:             uint(id),
		Identifier:     values[1],
		SpeciesId:      uint(speciesId),
		Height:         height,
		Weight:         weight,
		BaseExperience: baseExperience,
		Order:          order,
		IsDefault:      isDefault,
	}
	return t
}

func parsePokemonTypes(values []string) models.PokemonType {
	pokemonId, err := strconv.Atoi(values[0])
	if err != nil {
		log.Fatal("[pokemon][pokemonId] Unable to parse value as integer for "+values[0], err)
	}
	typeId, err := strconv.Atoi(values[1])
	if err != nil {
		log.Fatal("[pokemon][typeId] Unable to parse value as integer for "+values[1], err)
	}
	slot, err := strconv.Atoi(values[2])
	if err != nil {
		log.Fatal("[pokemon][slot] Unable to parse value as integer for "+values[2], err)
	}

	t := models.PokemonType{
		PokemonId: uint(pokemonId),
		TypeId:    uint(typeId),
		Slot:      slot,
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

	records = readCsvFile("data/pokemon_shapes.csv")
	for _, record := range records {
		shape := parseShape(record)
		db.Save(&shape)
	}

	records = readCsvFile("data/pokemon_colors.csv")
	for _, record := range records {
		color := parseColor(record)
		db.Save(&color)
	}

	records = readCsvFile("data/pokemon_species.csv")
	for _, record := range records {
		species := parseSpecies(record)
		db.Save(&species)
	}

	records = readCsvFile("data/pokemon.csv")
	for _, record := range records {
		pokemon := parsePokemon(record)
		db.Save(&pokemon)
	}

	records = readCsvFile("data/pokemon_types.csv")
	for _, record := range records {
		pokemonType := parsePokemonTypes(record)
		db.Clauses(clause.OnConflict{DoNothing: true}).Create(&pokemonType)
	}
}
