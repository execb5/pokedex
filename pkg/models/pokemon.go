package models

type Pokemon struct {
	ID             uint `gorm:"primary_key"`
	Identifier     string
	SpeciesId      uint
	Height         int
	Weight         int
	BaseExperience int
	Order          int
	IsDefault      bool
}

func (Pokemon) TableName() string {
	return "pokemon"
}
