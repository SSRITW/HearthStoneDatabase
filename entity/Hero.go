package entity

type Skill struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"not null"`
	ImageSrc string `gorm:"not null"`
	Expend int `gorm:"not null"`
	Describe string `gorm:"not null"`
}

type Hero struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT" form:"id" json:"id" binding:"required"`
	Name string `gorm:"not null" form:"name" json:"name" binding:"required"`
	ImageSrc string `gorm:"not null" form:"imageSrc" json:"imageSrc" binding:"required"`
	ProfessionId int `gorm:"not null" form:"professionId" json:"professionId" binding:"required"`
	SkillId int `gorm:"not null" form:"skillId" json:"skillId" binding:"required"`
}
