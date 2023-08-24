package models

type Pokemon struct {
	ID             uint
	Identifier     string
	SpeciesId      int64
	Height         int64
	Weight         int64
	BaseExperience int64
	Order          int64
	IsDefault      bool
}

func (Pokemon) TableName() string {
	return "pokemon"
}
