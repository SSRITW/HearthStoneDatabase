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
	/*gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return "t_" + defaultTableName;
	}*/

	var err error
	Db, err = gorm.Open(ConfigDataSource.Database.DriveName, ConfigDataSource.Database.ConnetionURL)
	Db.DB().SetMaxIdleConns(ConfigDataSource.Database.MaxIdleConns)
	Db.DB().SetMaxOpenConns(ConfigDataSource.Database.MaxOpenConns)
	if err !=nil {
		fmt.Println(err.Error())
	}
	Db.AutoMigrate(&entity.Profession{},&entity.Skill{},&entity.Hero{},&entity.CardType{},&entity.CardPackage{},&entity.CardBase{})
	Db.LogMode(true)
	return Db
}




