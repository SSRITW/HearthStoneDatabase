package service

import (
	"HearthStoneDatabase/entity"
	"HearthStoneDatabase/restgo"
)

/**
 * @arg pageSize
 * @arg pageNum 页码
 * @arg skill 查询条件
 * @return skills 技能数据切片
 */
func SkillsInfoPage(pageSize int,pageNum int,skill *entity.Skill)(skills []entity.Skill){
	restgo.Db.Where(&skill).Offset(restgo.GetPageOffset(pageSize,pageNum)).Limit(pageSize).Find(&skills)
	return
}

//获取该条件下的总条数
func SkillInfoCount(skill *entity.Skill)(count int){
	restgo.Db.Model(&skill).Where(&skill).Count(&count)
	return
}

/**
 * @arg pageSize
 * @arg id 技能id
 * @return skill
 */
func SkillInfoById(id int)(skill entity.Skill){
	restgo.Db.Where("id = ?",id).Find(&skill)
	return
}

//只更新struct里非空字段,返回影响条数
func SkillOfUpdate(skill *entity.Skill)int64{
	return restgo.Db.Model(&skill).Updates(&skill).RowsAffected
}

//插入一条新数据,返回影响条数
func SkillOfCreate(skill *entity.Skill)int64{
	return restgo.Db.Create(&skill).RowsAffected
}

func SkillOfDelete(id int)(int64){
	skill := entity.Skill{}
	skill.Id = id
	return restgo.Db.Delete(&skill).RowsAffected
}





