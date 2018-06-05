package entity

type CardType struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"not null"`
	Describe string `gorm:"not null"`
}

type CardPackage struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"not null"`
	BackImage string `gorm:"not null"`
}

type CardBase struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string `gorm:"not null"`
	Rarity string `gorm:"not null"`
	IsGolden string `gorm:"not null"`
	TypeId string `gorm:"not null"`
	IsNormal string `gorm:"not null"`
	Expend int `gorm:"not null"`
	Describe string `gorm:"not null"`
	CardPackageId int `gorm:"not null"`
	ProfessionId int `gorm:"not null"`
}
