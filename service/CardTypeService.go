package service

import (
	"github.com/jinzhu/gorm"
	"HearthStoneDatabase/entity"
	"HearthStoneDatabase/restgo"
)

/**
 * @arg pageSize
 * @arg pageNum 页码
 * @arg cardType 查询条件
 * @return cardTypes 卡牌类型数据切片
 */
func CardTypeInfoPage(db *gorm.DB,pageSize int,pageNum int,cardType *entity.CardType)(cardTypes []entity.CardType){
	db.Where(&cardType).Offset(restgo.GetPageOffset(pageSize,pageNum)).Limit(pageSize).Find(&cardTypes)
	return
}

//获取该条件下的数据总数
func CardTypeInfoCount(db *gorm.DB,cardType *entity.CardType)(count int){
	db.Model(&cardType).Where(&cardType).Count(&count)
	return
}

//通过id查询数据信息
func CardTypeInfoById(db *gorm.DB,id int)(cardType entity.CardType){
	db.Where("id = ?",id).Find(&cardType)
	return
}

//更新数据(仅更新有值部分
func CardTypeOfUpdate(db *gorm.DB,cardType *entity.CardType)(int64){
	return db.Model(&cardType).Updates(&cardType).RowsAffected
}

//插入一条数据
func CardTypeOfCreate(db *gorm.DB,cardType *entity.CardType)(int64){
	return db.Create(&cardType).RowsAffected
}
