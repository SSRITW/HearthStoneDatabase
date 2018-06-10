package service

import (
	"HearthStoneDatabase/entity"
	"HearthStoneDatabase/restgo"
)

/**
 * @arg pageSize
 * @arg pageNum 页码
 * @arg profession 查询条件
 * @return professions 技能数据切片
 */
func ProfessionInfoPage(pageSize int,pageNum int,profession *entity.Profession)(professions []entity.Profession){
	restgo.Db.Where(&profession).Offset(restgo.GetPageOffset(pageSize,pageNum)).Limit(pageSize).Find(&professions)
	return
}

//获取该条件下的数据数量
func ProfessionInfoCount(profession *entity.Profession)(count int){
	restgo.Db.Model(&profession).Where(&profession).Count(&count)
	return
}

//通过id查找数据
func ProfessionInfoById(id int)(profession entity.Profession){
	restgo.Db.Where("id = ?",id).Find(&profession)
	return
}
//仅更新struct内不为空的字段
func ProfessionOfUpdate(profession *entity.Profession)(int64){
	return restgo.Db.Model(&profession).Updates(&profession).RowsAffected
}

//插入一条数据
func ProfessionOfCreate(profession *entity.Profession)(int64){
	return restgo.Db.Create(&profession).RowsAffected
}

func ProfessionOfDelete(id int)(int64){
	profession := entity.Profession{}
	profession.Id = id

	return restgo.Db.Delete(&profession).RowsAffected
}
