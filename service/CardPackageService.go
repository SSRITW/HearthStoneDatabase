package service

import (
	"HearthStoneDatabase/entity"
	"HearthStoneDatabase/restgo"
)

//获取该条件下、该页码的数据
func CardPackageInfoPage(pageSize int,pageNum int,cardPackage *entity.CardPackage)(cardPackages []entity.CardPackage){
	restgo.Db.Where(&cardPackage).Offset(restgo.GetPageOffset(pageSize,pageNum)).Limit(pageSize).Find(&cardPackages)
	return
}

//获取该条件下的数据总数
func CardPackageInfoCount(cardPackage *entity.CardPackage)(count int){
	restgo.Db.Model(&cardPackage).Where(&cardPackage).Count(&count)
	return
}

//通过id获取该数据
func CardPackageInfoById(id int)(cardPackage entity.CardPackage){
	restgo.Db.Where("id = ?",id).Find(&cardPackage)
	return
}

//插入一条新数据
func CardPackageOfCreate(cardPackage *entity.CardPackage)(int64){
	return restgo.Db.Create(&cardPackage).RowsAffected
}

//更新数据
func CardPackageOfUpdate(cardPackage *entity.CardPackage)(int64){
	return restgo.Db.Model(&cardPackage).Updates(cardPackage).RowsAffected
}

func CardPackageOfDelete(id int)(int64){
	cardPackage := entity.CardPackage{}
	cardPackage.Id =  id
	return restgo.Db.Delete(&cardPackage).RowsAffected
}
