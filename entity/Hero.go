package entity

type Skill struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string
	ImageSrc string
}

type Hero struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"not null"`
	ImageSrc string `gorm:"not null"`
	ProfessionId int `gorm:"not null`
	SkillId int `gorm:"not null`
	Skill Skill `gorm:"ForeignKey:Id"`
}
