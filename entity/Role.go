package entity

type Role struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT" json:"id" form:"id" binding:"required"`
	Name string `gorm:"not null" json:"name" form:"name" binding:"required"`
	describe string `gorm:"not null" json:"describe" form:"describe" binding:"required"`
}
