package test

import (
	"testing"
	"HearthStoneDatabase/restgo"
	"HearthStoneDatabase/entity"
	"fmt"
	"HearthStoneDatabase/service"
)

func TestCreateSkill(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	skill := entity.Skill{0,"Crackle","skillImage3",2,"高级寒冰术"}
	flag := service.SkillOfCreate(&skill)
	fmt.Println(flag)
}

func TestGetSkillInfo(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	skill := entity.Skill{}
	skill.Expend = 2
	skills := service.SkillsInfoPage(1,1,&skill)
	for _,v:=range skills{
		fmt.Println(v)
	}
}

func TestGetSkillInfoCount(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	skill := entity.Skill{}
	skill.Expend = 2
	count := service.SkillInfoCount(&skill)
	fmt.Println(count)
}

func TestGetSkillInfoById(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	skill := service.SkillInfoById(1)
	fmt.Println(skill)
}

func TestUpdateSkill(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	skill := entity.Skill{}
	skill.Id = 1
	skill.Expend = 2
	flag := service.SkillOfUpdate(&skill)
	fmt.Println(flag)
}
