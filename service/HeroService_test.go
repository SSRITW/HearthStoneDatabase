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

	HeroCreate(db,&hero)

}


func TestGetHeroInfoPage(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	hero := entity.Hero{}
	hero.ProfessionId = 1
	heroes:= HeroInfoPage(db,1,1,&hero)
	for _,v:=range heroes {
		fmt.Println(v)
	}
}

func TestGetHeroInfoCount(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	hero := entity.Hero{}
	hero.ProfessionId = 1
	count:= HeroInfoCount(db,&hero)
	fmt.Println(count)
}

func TestGetHeroInfoById(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	hero := HeroInfoById(db,2)
	fmt.Println(hero)
}

func TestUpdateHero(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	hero := entity.Hero{1,"update2","updateSrc2",0,0}
	fmt.Println(HeroUpdate(db,&hero))
}

