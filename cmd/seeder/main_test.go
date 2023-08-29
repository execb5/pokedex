package main

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/execb5/pokedex/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestMain_parseType(t *testing.T) {
	tests := []struct {
		in   []string
		want models.Type
	}{
		{
			in: []string{
				"0",
				"bulba",
				"1",
				"2",
			},
			want: models.Type{
				ID:            uint(0),
				Identifier:    "bulba",
				GenerationId:  int64(1),
				DamageClassId: sql.NullInt64{Int64: int64(2), Valid: true},
			},
		},
		{
			in: []string{
				"3",
				"pikachu",
				"4",
				"",
			},
			want: models.Type{
				ID:            uint(3),
				Identifier:    "pikachu",
				GenerationId:  int64(4),
				DamageClassId: sql.NullInt64{Int64: 0, Valid: false},
			},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := parseType(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("models.Type(%s, %s, %s, %s)=%v; want %v", tt.in[0], tt.in[1], tt.in[2], tt.in[3], got, tt.want)
			}
		})
	}
}

func TestMain_parseShape(t *testing.T) {
	values := []string{"0", "round"}
	shape := parseShape(values)
	assert.Equal(t, shape.ID, uint(0), "they should be equal")
	assert.Equal(t, shape.Identifier, values[1], "they should be equal")
}

func TestMain_parseColor(t *testing.T) {
	values := []string{"0", "red"}
	color := parseColor(values)
	assert.Equal(t, color.ID, uint(0), "they should be equal")
	assert.Equal(t, color.Identifier, values[1], "they should be equal")
}

func TestMain_parseSpecies(t *testing.T) {
	tests := []struct {
		in   []string
		want models.Species
	}{
		{
			in: []string{"25", "pikachu", "1", "172", "10", "10", "8", "2", "4", "190", "50", "0", "10", "1", "2", "0", "0", "0", "26", "16"},
			want: models.Species{
				ID:                   uint(25),
				Identifier:           "pikachu",
				GenerationId:         1,
				EvolvesFromSpeciesId: sql.NullInt64{Int64: 172, Valid: true},
				EvolutionChainId:     10,
				ColorId:              uint(10),
				ShapeId:              sql.NullInt64{Int64: 8, Valid: true},
				HabitatId:            sql.NullInt64{Int64: 2, Valid: true},
				GenderRate:           4,
				CaptureRate:          190,
				BaseHappiness:        50,
				IsBaby:               false,
				HatchCounter:         10,
				HasGenderDifferences: true,
				GrowthRateId:         2,
				FormsSwitchable:      false,
				IsLegendary:          false,
				IsMythical:           false,
				Order:                26,
				ConquestOrder:        sql.NullInt64{Int64: 16, Valid: true},
			},
		},
		{
			in: []string{"25", "pikachu", "1", "", "10", "10", "8", "2", "4", "190", "50", "0", "10", "1", "2", "0", "0", "0", "26", "16"},
			want: models.Species{
				ID:                   uint(25),
				Identifier:           "pikachu",
				GenerationId:         1,
				EvolvesFromSpeciesId: sql.NullInt64{Int64: 0, Valid: false},
				EvolutionChainId:     10,
				ColorId:              uint(10),
				ShapeId:              sql.NullInt64{Int64: 8, Valid: true},
				HabitatId:            sql.NullInt64{Int64: 2, Valid: true},
				GenderRate:           4,
				CaptureRate:          190,
				BaseHappiness:        50,
				IsBaby:               false,
				HatchCounter:         10,
				HasGenderDifferences: true,
				GrowthRateId:         2,
				FormsSwitchable:      false,
				IsLegendary:          false,
				IsMythical:           false,
				Order:                26,
				ConquestOrder:        sql.NullInt64{Int64: 16, Valid: true},
			},
		},
		{
			in: []string{"25", "pikachu", "1", "172", "10", "10", "", "2", "4", "190", "50", "0", "10", "1", "2", "0", "0", "0", "26", "16"},
			want: models.Species{
				ID:                   uint(25),
				Identifier:           "pikachu",
				GenerationId:         1,
				EvolvesFromSpeciesId: sql.NullInt64{Int64: 172, Valid: true},
				EvolutionChainId:     10,
				ColorId:              uint(10),
				ShapeId:              sql.NullInt64{Int64: 0, Valid: false},
				HabitatId:            sql.NullInt64{Int64: 2, Valid: true},
				GenderRate:           4,
				CaptureRate:          190,
				BaseHappiness:        50,
				IsBaby:               false,
				HatchCounter:         10,
				HasGenderDifferences: true,
				GrowthRateId:         2,
				FormsSwitchable:      false,
				IsLegendary:          false,
				IsMythical:           false,
				Order:                26,
				ConquestOrder:        sql.NullInt64{Int64: 16, Valid: true},
			},
		},
		{
			in: []string{"25", "pikachu", "1", "172", "10", "10", "8", "", "4", "190", "50", "0", "10", "1", "2", "0", "0", "0", "26", "16"},
			want: models.Species{
				ID:                   uint(25),
				Identifier:           "pikachu",
				GenerationId:         1,
				EvolvesFromSpeciesId: sql.NullInt64{Int64: 172, Valid: true},
				EvolutionChainId:     10,
				ColorId:              uint(10),
				ShapeId:              sql.NullInt64{Int64: 8, Valid: true},
				HabitatId:            sql.NullInt64{Int64: 0, Valid: false},
				GenderRate:           4,
				CaptureRate:          190,
				BaseHappiness:        50,
				IsBaby:               false,
				HatchCounter:         10,
				HasGenderDifferences: true,
				GrowthRateId:         2,
				FormsSwitchable:      false,
				IsLegendary:          false,
				IsMythical:           false,
				Order:                26,
				ConquestOrder:        sql.NullInt64{Int64: 16, Valid: true},
			},
		},
		{
			in: []string{"25", "pikachu", "1", "172", "10", "10", "8", "2", "4", "190", "50", "0", "10", "1", "2", "0", "0", "0", "26", ""},
			want: models.Species{
				ID:                   uint(25),
				Identifier:           "pikachu",
				GenerationId:         1,
				EvolvesFromSpeciesId: sql.NullInt64{Int64: 172, Valid: true},
				EvolutionChainId:     10,
				ColorId:              uint(10),
				ShapeId:              sql.NullInt64{Int64: 8, Valid: true},
				HabitatId:            sql.NullInt64{Int64: 2, Valid: true},
				GenderRate:           4,
				CaptureRate:          190,
				BaseHappiness:        50,
				IsBaby:               false,
				HatchCounter:         10,
				HasGenderDifferences: true,
				GrowthRateId:         2,
				FormsSwitchable:      false,
				IsLegendary:          false,
				IsMythical:           false,
				Order:                26,
				ConquestOrder:        sql.NullInt64{Int64: 0, Valid: false},
			},
		},
		{
			in: []string{"25", "pikachu", "1", "", "10", "10", "", "", "4", "190", "50", "0", "10", "1", "2", "0", "0", "0", "26", ""},
			want: models.Species{
				ID:                   uint(25),
				Identifier:           "pikachu",
				GenerationId:         1,
				EvolvesFromSpeciesId: sql.NullInt64{Int64: 0, Valid: false},
				EvolutionChainId:     10,
				ColorId:              uint(10),
				ShapeId:              sql.NullInt64{Int64: 0, Valid: false},
				HabitatId:            sql.NullInt64{Int64: 0, Valid: false},
				GenderRate:           4,
				CaptureRate:          190,
				BaseHappiness:        50,
				IsBaby:               false,
				HatchCounter:         10,
				HasGenderDifferences: true,
				GrowthRateId:         2,
				FormsSwitchable:      false,
				IsLegendary:          false,
				IsMythical:           false,
				Order:                26,
				ConquestOrder:        sql.NullInt64{Int64: 0, Valid: false},
			},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := parseSpecies(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("models.Species(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)=%v; want %v", tt.in[0], tt.in[1], tt.in[2], tt.in[3], tt.in[4], tt.in[5], tt.in[6], tt.in[7], tt.in[8], tt.in[9], tt.in[10], tt.in[11], tt.in[12], tt.in[13], tt.in[14], tt.in[15], tt.in[16], tt.in[17], tt.in[18], tt.in[19], got, tt.want)
			}
		})
	}
}
