package entity

type Skill struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT" json:"id" form:"id" binding:"required" `
	Name string `gorm:"not null" json:"name" form:"name" binding:"required" `
	ImageSrc string `gorm:"not null" json:"image_src" form:"image_src" binding:"required"`
	Expend int `gorm:"not null" json:"expend" form:"expend" binding:"required"`
	Describe string `gorm:"not null" json:"describe" form:"describe" binding:"required"`
}

type Hero struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT" form:"id" json:"id" binding:"required"`
	Name string `gorm:"not null" form:"name" json:"name" binding:"required"`
	ImageSrc string `gorm:"not null" form:"imageSrc" json:"imageSrc" binding:"required"`
	ProfessionId int `gorm:"not null" form:"professionId" json:"professionId" binding:"required"`
	SkillId int `gorm:"not null" form:"skillId" json:"skillId" binding:"required"`
}
