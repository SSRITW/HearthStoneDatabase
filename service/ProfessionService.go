package service

import (
	"github.com/jinzhu/gorm"
	"HearthStoneDatabase/entity"
	"HearthStoneDatabase/restgo"
)

/**
 * @arg pageSize
 * @arg pageNum 页码
 * @arg profession 查询条件
 * @return professions 技能数据切片
 */
func ProfessionInfoPage(db *gorm.DB,pageSize int,pageNum int,profession *entity.Profession)(professions []entity.Profession){
	db.Where(&profession).Offset(restgo.GetPageOffset(pageSize,pageNum)).Limit(pageSize).Find(&professions)
	return
}

//获取该条件下的数据数量
func ProfessionInfoCount(db *gorm.DB,profession *entity.Profession)(count int){
	db.Model(&profession).Where(&profession).Count(&count)
	return
}

//通过id查找数据
func ProfessionInfoById(db *gorm.DB,id int)(profession entity.Profession){
	db.Where("id = ?",id).Find(&profession)
	return
}
//仅更新struct内不为空的字段
func ProfessionOfUpdate(db *gorm.DB,profession *entity.Profession)(int64){
	return db.Model(&profession).Updates(&profession).RowsAffected
}

//插入一条数据
func ProfessionOfCreate(db *gorm.DB,profession *entity.Profession)(int64){
	return db.Create(&profession).RowsAffected
}
