package restgo

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"HearthStoneDatabase/entity"
)

var Db *gorm.DB

func OpenDBConnect()(*gorm.DB){

	//更改默认表名
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return "f_" + defaultTableName;
	}

	var err error
	Db, err = gorm.Open("mysql", "root:123456@(localhost:3306)/hearthstonedb?charset=utf8&parseTime=True&loc=Local")
	initRelated(Db)
	if err !=nil {
		fmt.Println(err.Error())
	}
	Db.AutoMigrate(&entity.Profession{},&entity.Skill{},&entity.Hero{},&entity.CardType{},&entity.CardPackage{},&entity.CardBase{})
	return Db
}

func initRelated(db *gorm.DB){
	db.Model(&entity.Hero{}).Related(&entity.Skill{})
}


