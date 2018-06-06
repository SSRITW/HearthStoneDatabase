package service

import (
	"testing"
	"HearthStoneDatabase/restgo"
	"fmt"
	"HearthStoneDatabase/entity"
)

func TestAddHero(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()

	hero := entity.Hero{}
	hero.Name = "test3"
	hero.ImageSrc = "testImage3"
	hero.ProfessionId = 2
	hero.SkillId = 1

	AddHero(db,&hero)

}


func TestGetHeroInfo(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	hero := entity.Hero{}
	hero.ProfessionId = 1
	heroes := GetHeroInfo(db,1,2,&hero)
	for _,v:=range heroes {
		fmt.Println(v)
	}
}


func TestGetHeroInfoById(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	hero := GetHeroInfoById(db,2)
	fmt.Println(hero)
}

func TestUpdateHero(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	hero := entity.Hero{1,"update","updateSrc",1,1}
	UpdateHero(db,&hero)
}

