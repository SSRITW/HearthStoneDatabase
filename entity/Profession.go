package entity

type Profession struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string	`gorm:"not null"`
}

