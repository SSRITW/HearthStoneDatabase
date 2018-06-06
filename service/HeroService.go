package service

import (
	"HearthStoneDatabase/entity"
	"github.com/jinzhu/gorm"
	"HearthStoneDatabase/model"
	"HearthStoneDatabase/restgo"
)

const
(
	_TABLE_NAME = `t_heros hero`
	_SELECT_STR = `hero.*,skill.name as skill_name,skill.image_src as skill_image_src,profession.name as profession_name`
	_JOIN_SKILL = `join t_skills skill on skill.id = hero.skill_id`
	_JOIN_PROFESSION = `join t_professions profession on profession.id = hero.profession_id`
)

func GetHeroInfo(db *gorm.DB,pageSize int,pageNum int,hero *entity.Hero)(heroes []model.Hero){
	getWhereQuery(db,hero).Select(_SELECT_STR).Joins(_JOIN_SKILL).Joins(_JOIN_PROFESSION).Table(_TABLE_NAME).
		Offset(restgo.GetPageOffset(pageSize,pageNum)).Limit(pageSize).Scan(&heroes)
	return
}

func GetHeroInfoById(db *gorm.DB,id int)(hero model.Hero){
	/*db.Debug().Where("id = ?",id).Find(&hero)
	db.Debug().Model(&hero).Related(&hero.Skill)*/

	db.Table(_TABLE_NAME).Select(_SELECT_STR).
		Joins(_JOIN_SKILL).Joins(_JOIN_PROFESSION).Where("hero.id = ?",id).Scan(&hero)
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
		db = db.Where("hero.name = ?",hero.Name)
	}

	if hero.ProfessionId!=0 {
		db = db.Where("hero.profession_id = ?",hero.ProfessionId)
	}

	return db

}