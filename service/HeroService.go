package service

import (
	"HearthStoneDatabase/entity"
	"github.com/jinzhu/gorm"
	"HearthStoneDatabase/model"
)

const
(
	tableName = `t_heros hero`
	selectStr = `hero.*,skill.name as skill_name,skill.image_src as skill_image_src,profession.name as profession_name`
	joinSkill = `join t_skills skill on skill.id = hero.skill_id`
	joinProfession = `join t_professions profession on profession.id = hero.profession_id`
)

func GetAllHeroInfo(db *gorm.DB)(heroes []model.Hero){
	db.Debug().Table("t_heros hero")
	return
}

func GetHeroInfo(db *gorm.DB,hero *entity.Hero)(heroes []model.Hero){
	getWhereQuery(db,hero).Select(selectStr).Joins(joinSkill).Joins(joinProfession).Table(tableName).Scan(&heroes)
	return
}

func GetHeroInfoById(db *gorm.DB,id int)(hero model.Hero){
	/*db.Debug().Where("id = ?",id).Find(&hero)
	db.Debug().Model(&hero).Related(&hero.Skill)*/

	db.Table(tableName).Select(selectStr).
		Joins(joinSkill).Joins(joinProfession).Where("hero.id = ?",id).Scan(&hero)
	return
}

func AddHero(db *gorm.DB,hero *entity.Hero){
	db.Create(&hero)
}

func UpdateHero(db *gorm.DB,hero *entity.Hero){
	db.Save(&hero)
}

func getWhereQuery(db *gorm.DB,hero *entity.Hero)(*gorm.DB){

	if hero.Name!="" {
		db.Where("hero.name = ?",hero.Name)
	}

	if hero.ProfessionId!=0 {
		db.Where("hero.profession_id = ?",hero.ProfessionId)
	}

	return db

}