package service

import (
	"github.com/jinzhu/gorm"
	"HearthStoneDatabase/entity"
	"HearthStoneDatabase/restgo"
)

/**
 * @arg pageSize
 * @arg pageNum 页码
 * @arg hero 查询条件
 * @return skills 技能数据切片
 */
func SkillsInfoPage(db *gorm.DB,pageSize int,pageNum int,skill *entity.Skill)(skills []entity.Skill){
	db.Where(&skill).Offset(restgo.GetPageOffset(pageSize,pageNum)).Limit(pageSize).Find(&skills)
	return
}

//获取该条件下的总条数
func SkillInfosCount(db *gorm.DB,skill *entity.Skill)(count int){
	db.Model(&skill).Where(&skill).Count(&count)
	return
}

/**
 * @arg pageSize
 * @arg id 技能id
 * @return skill
 */
func SkillInfoById(db *gorm.DB,id int)(skill *entity.Skill){
	db.Where("id = ?",id).Find(&skill)
	return
}

//只更新struct里非空字段,返回影响条数
func SkillUpdate(db *gorm.DB,skill *entity.Skill)int64{
	return db.Model(&skill).Updates(&skill).RowsAffected
}

//插入一条新数据,返回影响条数
func SkillCreate(db *gorm.DB,skill *entity.Skill)int64{
	return db.Create(&skill).RowsAffected
}





