package test

import (
	"testing"
	"HearthStoneDatabase/restgo"
	"HearthStoneDatabase/entity"
	"fmt"
	"HearthStoneDatabase/service"
)

func TestCardTypeOfCreate(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()

	cardType := entity.CardType{0,"随从","召唤一个随从"}
	flag := service.CardTypeOfCreate(&cardType)
	fmt.Println(flag)
}

func TestCardTypeOfUpdate(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()

	cardType := entity.CardType{2,"","施法一个法术"}
	flag := service.CardTypeOfUpdate(&cardType)
	fmt.Println(flag)
}

func TestCardTypeInfoById(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	cardType := service.CardTypeInfoById(1)
	fmt.Println(cardType)
}

func TestCardTypeInfoCount(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	cardType := entity.CardType{}
	count := service.CardTypeInfoCount(&cardType)
	fmt.Println(count)
}

func TestCardTypeInfoPage(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	cardType := entity.CardType{0, "", ""}
	cardTypes := service.CardTypeInfoPage(1,2,&cardType)
	for _,v:=range cardTypes{
		fmt.Println(v)
	}
}