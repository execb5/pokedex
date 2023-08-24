package models

import "fmt"

type Pokemon struct {
	ID             int64
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

// create table pokemon
// (
//     id              serial,
//     identifier      varchar(100) not null,
//     species_id      integer      not null,
//     height          integer      not null,
//     weight          integer      not null,
//     base_experience integer      not null,
//     "order"         integer      not null,
//     is_default      boolean      not null
// );

func Seed() {
	a := Pokemon{}
	fmt.Println(a)
}
