package service

import (
	"github.com/jinzhu/gorm"
	"HearthStoneDatabase/entity"
	"HearthStoneDatabase/restgo"
)

func CardPackageInfoPage(db *gorm.DB,pageSize int,pageNum int,cardPackage *entity.CardPackage)(cardPackages []entity.CardPackage){
	db.Where(&cardPackage).Offset(restgo.GetPageOffset(pageSize,pageNum)).Limit(pageSize).Find(&cardPackages)
	return
}

func CardPackageInfoCount(db *gorm.DB,cardPackage *entity.CardPackage)(count int){
	db.Model(&cardPackage).Where(&cardPackage).Count(&count)
	return
}

func CardPackageInfoById(db *gorm.DB,id int)(cardPackage entity.CardPackage){
	db.Where("id = ?",id).Find(&cardPackage)
	return
}

func CardPackageOfCreate(db *gorm.DB,cardPackage *entity.CardPackage)(int64){
	return db.Create(&cardPackage).RowsAffected
}

func CardPackageOfUpdate(db *gorm.DB,cardPackage *entity.CardPackage)(int64){
	return db.Model(&cardPackage).Updates(cardPackage).RowsAffected
}
