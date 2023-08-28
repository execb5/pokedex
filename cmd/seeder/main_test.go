package main

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/execb5/pokedex/pkg/models"
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
