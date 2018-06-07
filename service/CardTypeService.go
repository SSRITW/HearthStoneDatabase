package service

import (
	"HearthStoneDatabase/entity"
	"HearthStoneDatabase/restgo"
)

/**
 * @arg pageSize
 * @arg pageNum 页码
 * @arg cardType 查询条件
 * @return cardTypes 卡牌类型数据切片
 */
func CardTypeInfoPage(pageSize int,pageNum int,cardType *entity.CardType)(cardTypes []entity.CardType){
	restgo.Db.Where(&cardType).Offset(restgo.GetPageOffset(pageSize,pageNum)).Limit(pageSize).Find(&cardTypes)
	return
}

//获取该条件下的数据总数
func CardTypeInfoCount(cardType *entity.CardType)(count int){
	restgo.Db.Model(&cardType).Where(&cardType).Count(&count)
	return
}

//通过id查询数据信息
func CardTypeInfoById(id int)(cardType entity.CardType){
	restgo.Db.Where("id = ?",id).Find(&cardType)
	return
}

//更新数据(仅更新有值部分
func CardTypeOfUpdate(cardType *entity.CardType)(int64){
	return restgo.Db.Model(&cardType).Updates(&cardType).RowsAffected
}

//插入一条数据
func CardTypeOfCreate(cardType *entity.CardType)(int64){
	return restgo.Db.Create(&cardType).RowsAffected
}
