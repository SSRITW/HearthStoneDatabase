package entity

type Profession struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT" json:"id" form:"id" binding:"required"`
	Name string	`gorm:"not null" json:"name" form:"name" binding:"required"`
	CardFront string `gorm:"not null" json:"card_front" form:"card_front" binding:"required"`
}

