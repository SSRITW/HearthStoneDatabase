package model

type User struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT"`
	Nickname string `gorm:"not null"`
	Sex int `gorm:"not null"`
	LoginName string `gorm:"not null"`
	Password string `gorm:"not null"`
}
