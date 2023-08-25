package models

import (
	"database/sql"
)

type Species struct {
	ID                   uint `gorm:"primary_key"`
	Identifier           string
	GenerationId         int
	EvolvesFromSpeciesId sql.NullInt64
	EvolutionChainId     int
	ColorId              uint
	ShapeId              sql.NullInt64
	HabitatId            sql.NullInt64
	GenderRate           int
	CaptureRate          int
	BaseHappiness        int
	IsBaby               bool
	HatchCounter         int
	HasGenderDifferences bool
	GrowthRateId         int
	FormsSwitchable      bool
	IsLegendary          bool
	IsMythical           bool
	Order                int
	ConquestOrder        sql.NullInt64
}
