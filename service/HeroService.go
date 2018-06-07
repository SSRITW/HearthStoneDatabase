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
	_SELECT_STR = `hero.*,skill.name as skill_name,skill.image_src as skill_image_src,skill.expend as skill_expend,skill.describe as skill_describe,profession.name as profession_name`
	_JOIN_SKILL = `join t_skills skill on skill.id = hero.skill_id`
	_JOIN_PROFESSION = `join t_professions profession on profession.id = hero.profession_id`
)

/**
 * @arg pageSize
 * @arg pageNum 页码
 * @arg hero 查询条件
 * @return heroes 英雄数据切片
 */
func HeroInfoPage(pageSize int,pageNum int,hero *entity.Hero)(heroes []model.Hero){
	getWhereQuery(restgo.Db,hero).Select(_SELECT_STR).Joins(_JOIN_SKILL).Joins(_JOIN_PROFESSION).Table(_TABLE_NAME).
		Offset(restgo.GetPageOffset(pageSize,pageNum)).Limit(pageSize).Scan(&heroes)
	return
}

//获取该条件下的总条数
func HeroInfoCount(hero *entity.Hero)(count int){
	getWhereQuery(restgo.Db,hero).Joins(_JOIN_SKILL).Joins(_JOIN_PROFESSION).Table(_TABLE_NAME).Count(&count)
	return
}

/**
 * @arg pageSize
 * @arg id 英雄id
 * @return hero
 */
func HeroInfoById(id int)(hero model.Hero){
	/*db.Debug().Where("id = ?",id).Find(&hero)
	db.Debug().Model(&hero).Related(&hero.Skill)*/

	restgo.Db.Table(_TABLE_NAME).Select(_SELECT_STR).
		Joins(_JOIN_SKILL).Joins(_JOIN_PROFESSION).Where("hero.id = ?",id).Scan(&hero)
	return
}

//插入一条新数据,返回影响条数
func HeroOfCreate(hero *entity.Hero)(int64){
	return restgo.Db.Create(&hero).RowsAffected
}

//只更新struct里非空字段,返回影响条数
func HeroOfUpdate(hero *entity.Hero)(int64){
	return restgo.Db.Model(&hero).Updates(&hero).RowsAffected
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