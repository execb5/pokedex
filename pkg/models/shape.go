package models

type Shape struct {
	ID         uint `gorm:"primary_key"`
	Identifier string
}
