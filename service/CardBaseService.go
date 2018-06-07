package service

import (
	"HearthStoneDatabase/entity"
	"HearthStoneDatabase/model"
	"HearthStoneDatabase/restgo"
	"github.com/jinzhu/gorm"
)

const (
	_CARD_TABLE_NAME = "t_card_bases cBase"
	_CARD_JOIN_PACKAGE = "JOIN t_card_packages cPackage on cBase.card_package_id = cPackage.id"
	_CARD_JOIN_TYPE = "JOIN t_card_types cType on cBase.type_id = cType.id"
	_CARD_JOIN_PROFESSION = "JOIN t_professions profession on cBase.profession_id = profession.id"
	_CARD_SELECT_STR = "cBase.*, cType.name as card_type_name, cPackage.name as card_package_name," +
		" cPackage.back_image as card_package_back_image, profession.card_front as profession_card_front"
)

func CardBaseInfoPage(pageSize int, pageNum int, base *entity.CardBase)(bases []model.CardBase){
	getCardWhereQuery(restgo.Db,base).Table(_CARD_TABLE_NAME).Select(_CARD_SELECT_STR).Joins(_CARD_JOIN_PACKAGE).Joins(_CARD_JOIN_PROFESSION).
		Joins(_CARD_JOIN_TYPE).Offset(restgo.GetPageOffset(pageSize,pageNum)).Limit(pageSize).Scan(&bases)
	return
}

func CardBaseInfoCount(base *entity.CardBase)(count int){
	getCardWhereQuery(restgo.Db,base).Table(_CARD_TABLE_NAME).Joins(_CARD_JOIN_PACKAGE).
		Joins(_CARD_JOIN_PROFESSION).Joins(_CARD_JOIN_TYPE).Count(&count)
	return
}

func CardBaseInfoById(id int)(base model.CardBase){
	restgo.Db.Table(_CARD_TABLE_NAME).Select(_CARD_SELECT_STR).Joins(_CARD_JOIN_PACKAGE).Joins(_CARD_JOIN_PROFESSION).Joins(_CARD_JOIN_TYPE).
		Where("cBase.id = ?",id).Scan(&base)
	return
}

func CardBaseOfUpdate(base *entity.CardBase)(int64){
	return restgo.Db.Model(&base).Updates(&base).RowsAffected
}

func CardBseOfCreate(base *entity.CardBase)(int64){
	return restgo.Db.Create(&base).RowsAffected
}

func getCardWhereQuery(db *gorm.DB,base *entity.CardBase)(*gorm.DB){

	if base.Name != "" {
		db = db.Where("cBase.name like '%"+base.Name+"%'")
	}

	if base.Rarity != "" {
		db = db.Where("cBase.rarity = ?" , base.Rarity)
	}

	if base.IsGolden != "" {
		db = db.Where("cBase.is_golden = ?" , base.IsGolden)
	}

	if base.TypeId !=0 {
		db = db.Where("cBase.type_id = ?" , base.TypeId)
	}

	if base.IsNormal != "" {
		db = db.Where("cBase.is_normal = ?",base.IsNormal)
	}

	if base.Expend != 0 {
		db = db.Where("cBase.expend = ?",base.Expend)
	}

	if base.CardPackageId != 0 {
		db = db.Where("cBase.card_package_id = ?",base.CardPackageId)
	}

	if base.TypeId != 0 {
		db = db.Where("cBase.type_id = ?",base.TypeId)
	}
	
	return db
}




