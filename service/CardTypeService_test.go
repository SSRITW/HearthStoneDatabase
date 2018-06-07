package service

import (
	"testing"
	"HearthStoneDatabase/restgo"
	"HearthStoneDatabase/entity"
	"fmt"
)

func TestCardTypeOfCreate(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()

	cardType := entity.CardType{0,"法术","施法一个法术"}
	flag := CardTypeOfCreate(db,&cardType)
	fmt.Println(flag)
}

func TestCardTypeOfUpdate(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()

	cardType := entity.CardType{2,"","施法一个法术"}
	flag := CardTypeOfUpdate(db,&cardType)
	fmt.Println(flag)
}

func TestCardTypeInfoById(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	cardType := CardTypeInfoById(db,1)
	fmt.Println(cardType)
}

func TestCardTypeInfoCount(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	cardType := entity.CardType{}
	count := CardTypeInfoCount(db,&cardType)
	fmt.Println(count)
}

func TestCardTypeInfoPage(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	cardType := entity.CardType{0, "", ""}
	cardTypes := CardTypeInfoPage(db,1,2,&cardType)
	for _,v:=range cardTypes{
		fmt.Println(v)
	}
}