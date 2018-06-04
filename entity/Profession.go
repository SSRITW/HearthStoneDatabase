package entity

import "github.com/jinzhu/gorm"

type Professtion struct {
	gorm.Model
	Id int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string	`gorm:"not null"`
}

