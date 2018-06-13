package model

type Menu struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT"`
	ParentId int `gorm:"not null"`
	Name string `gorm:"not null"`
	Url string `gorm:"not null"`
	Action string `gorm:"not null"`
}