package models

type Color struct {
	ID         uint `gorm:"primary_key"`
	Identifier string
}
