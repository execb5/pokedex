package models

type Pokemon struct {
	ID             uint   `gorm:"primary_key" json:"id"`
	Identifier     string `json:"identifier"`
	SpeciesId      uint   `json:"species_id"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	BaseExperience int    `json:"base_experience"`
	Order          int    `json:"order"`
	IsDefault      bool   `json:"is_default"`
}

func (Pokemon) TableName() string {
	return "pokemon"
}
