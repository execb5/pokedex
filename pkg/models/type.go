package models

import "database/sql"

type Type struct {
	ID            uint `gorm:"primary_key"`
	Identifier    string
	GenerationId  int64
	DamageClassId sql.NullInt64
}
