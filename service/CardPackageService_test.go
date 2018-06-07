package service

import (
	"testing"
	"HearthStoneDatabase/restgo"
	"HearthStoneDatabase/entity"
	"fmt"
)

func TestCardPackageOfCreate(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()

	cardPackage := entity.CardPackage{0,"狗头人","狗头人卡背"}
	count := CardPackageOfCreate(db,&cardPackage)
	fmt.Println(count)
}

func TestCardPackageOfUpdate(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()

	cardPackage := entity.CardPackage{0,"","女巫森林卡背"}
	count := CardPackageOfUpdate(db,&cardPackage)
	fmt.Println(count)
}

func TestCardPackageInfoById(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()

	cardPackage := CardPackageInfoById(db,2)
	fmt.Println(cardPackage)
}

func TestCardPackageInfoCount(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	cardPackage := entity.CardPackage{}
	count := CardPackageInfoCount(db,&cardPackage)
	fmt.Println(count)
}

func TestCardPackageInfoPage(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()

	cardPackage := entity.CardPackage{}

	packages := CardPackageInfoPage(db,2,1,&cardPackage)
	fmt.Println(packages)
}
