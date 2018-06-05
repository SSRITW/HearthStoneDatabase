package service

import (
	"HearthStoneDatabase/entity"
	"github.com/jinzhu/gorm"
)

func GetAllHeroInfo(db *gorm.DB)(heroes []entity.Hero){
	var hero entity.Hero
	db.Debug().Model(&hero).Find(&heroes).Related(&hero.Skill)
	return
}

func GetHeroInfo(db *gorm.DB,hero *entity.Hero)(heroes []entity.Hero){
	db.Where(&hero).Find(&heroes)
	return
}

func GetHeroInfoById(db *gorm.DB,id int)(hero entity.Hero){
	db.Debug().Where("id = ?",id).Find(&hero).Related(&hero.Skill)
	return
}

func AddHero(db *gorm.DB,hero *entity.Hero){
	db.Create(&hero)
}

func UpdateHero(db *gorm.DB,hero *entity.Hero){
	db.Save(&hero)
}