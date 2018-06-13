package model

type Role struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"not null"`
	describe string `gorm:"not null"`
}

type RolePower struct{
	Id int `gorm:"primary_key;AUTO_INCREMENT"`
	RoleId int `gorm:"not null"`
	MenuId int `gorm:"not null"`
}
