package entity

import "github.com/jinzhu/gorm"

type Skill struct {
	gorm.Model
	Id int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"not null"`
	ImageSrc string `gorm:"not null;"`
}

type Hero struct {
	gorm.Model
	Id int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"not null"`
	ImageSrc string `gorm:"not null"`
	Professtion
	Skill
}
