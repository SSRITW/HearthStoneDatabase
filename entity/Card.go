package entity

type CardType struct {
	Id int `json:"id" form:"id" gorm:"primary_key;AUTO_INCREMENT" binding:"required"`
	Name string `json:"name" form:"name" gorm:"not null" binding:"required"`
	Describe string `json:"describe" form:"describe" gorm:"not null" binding:"required"`
}

type CardPackage struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT" json:"id" form:"id" binding:"required"`
	Name string `gorm:"not null" json:"name" form:"name" binding:"required"`
	BackImage string `gorm:"not null" json:"back_image" form:"back_image" binding:"required"`
}

type CardBase struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT" json:"id" form:"id" binding:"required"`
	Name string `gorm:"not null" json:"name" form:"name" binding:"required"`
	Rarity string `gorm:"not null" json:"rarity" form:"rarity" binding:"required"`
	IsGolden string `gorm:"not null" json:"is_golden" form:"is_golden" binding:"required"`
	TypeId int `gorm:"not null" json:"type_id" form:"type_id" binding:"required"`
	IsNormal string `gorm:"not null" json:"is_normal" form:"is_normal" binding:"required"`
	Expend int `gorm:"not null" json:"expend" form:"expend" binding:"required"`
	Image string `gorm:"not null" json:"image" form:"image" binding:"required"`
	Describe string `gorm:"not null" json:"describe" form:"describe" binding:"required"`
	CardPackageId int `gorm:"not null" json:"card_package_id" form:"card_package_id" binding:"required"`
	ProfessionId int `gorm:"not null" json:"profession_id" form:"profession_id" binding:"required"`
}
