package entity

type User struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT" json:"id" form:"id" binding:"required"`
	Nickname string `gorm:"not null" json:"nickname" form:"nickname" binding:"required"`
	Sex int `gorm:"not null" json:"sex" form:"sex" binding:"required"`
	LoginName string `gorm:"not null" json:"login_name" form:"login_name" binding:"required"`
	Password string `gorm:"not null" json:"password" form:"password" binding:"required"`
	Salt string `gorm:"not null" json:"salt" form:"salt" binding:"required"`
	RoleId int `gorm:"not null" json:"role_id" form:"role_id" binding:"required"`
}
