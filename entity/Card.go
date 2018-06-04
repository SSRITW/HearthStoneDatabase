package entity

import "github.com/jinzhu/gorm"

type CardType struct {
	gorm.Model
	Id int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"not null"`
	Describe string `gorm:"not null"`
}

type CardPackage struct {
	gorm.Model
	Id int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"not null"`
	BackImage string `gorm:"not null"`
}

type CardBase struct {
	gorm.Model
	Id int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"not null"`
	Rarity string `gorm:"not null"`
	IsGolden string `gorm:"not null"`
	TypeId string `gorm:"not null"`
	IsNormal string `gorm:"not null"`
	Expend int `gorm:"not null"`
	Describe string `gorm:"not null"`
	CardPackage
	Professtion
}
