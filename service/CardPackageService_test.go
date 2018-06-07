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

	cardPackage := entity.CardPackage{0,"冰封王座","冰封王座卡背"}
	count := CardPackageOfCreate(&cardPackage)
	fmt.Println(count)
}

func TestCardPackageOfUpdate(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()

	cardPackage := entity.CardPackage{1,"","女巫森林卡背-"}
	count := CardPackageOfUpdate(&cardPackage)
	fmt.Println(count)
}

func TestCardPackageInfoById(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()

	cardPackage := CardPackageInfoById(2)
	fmt.Println(cardPackage)
}

func TestCardPackageInfoCount(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	cardPackage := entity.CardPackage{}
	count := CardPackageInfoCount(&cardPackage)
	fmt.Println(count)
}

func TestCardPackageInfoPage(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()

	cardPackage := entity.CardPackage{}

	packages := CardPackageInfoPage(2,1,&cardPackage)
	fmt.Println(packages)
}
