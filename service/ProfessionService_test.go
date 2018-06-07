package service

import (
	"testing"
	"HearthStoneDatabase/restgo"
	"HearthStoneDatabase/entity"
	"fmt"
)

func TestProfessionOfCreate(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	profession := entity.Profession{0,"牧师"}
	flag := ProfessionOfCreate(&profession)
	fmt.Println(flag)
}

func TestProfessionOfUpdate(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	profession := entity.Profession{1,"法师"}
	flag := ProfessionOfUpdate(&profession)
	fmt.Println(flag)
}

func TestProfessionInfoById(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	profession := ProfessionInfoById(2)
	fmt.Println(profession)
}

func TestProfessionInfoCount(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	profession := entity.Profession{0,"法师"}
	count := ProfessionInfoCount(&profession);
	fmt.Println(count)
}

func TestProfessionInfoPage(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	//profession := entity.Profession{0,"牧师"}
	profession := entity.Profession{}
	result := ProfessionInfoPage(5,1,&profession);
	for _,v:=range result{
		fmt.Println(v)
	}
}