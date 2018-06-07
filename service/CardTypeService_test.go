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

	cardType := entity.CardType{0,"随从","召唤一个随从"}
	flag := CardTypeOfCreate(&cardType)
	fmt.Println(flag)
}

func TestCardTypeOfUpdate(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()

	cardType := entity.CardType{2,"","施法一个法术"}
	flag := CardTypeOfUpdate(&cardType)
	fmt.Println(flag)
}

func TestCardTypeInfoById(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	cardType := CardTypeInfoById(1)
	fmt.Println(cardType)
}

func TestCardTypeInfoCount(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	cardType := entity.CardType{}
	count := CardTypeInfoCount(&cardType)
	fmt.Println(count)
}

func TestCardTypeInfoPage(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	cardType := entity.CardType{0, "", ""}
	cardTypes := CardTypeInfoPage(1,2,&cardType)
	for _,v:=range cardTypes{
		fmt.Println(v)
	}
}