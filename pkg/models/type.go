package models

type Type struct {
	ID            uint `gorm:"primary_key"`
	Identifier    string
	GenerationId  int64
	DamageClassId int64
}
