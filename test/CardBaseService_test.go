package test

import (
	"testing"
	"HearthStoneDatabase/restgo"
	"HearthStoneDatabase/entity"
	"fmt"
	"HearthStoneDatabase/service"
	"strings"
)

func TestCardBseOfCreate(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()

	card := entity.CardBase{0,"法术浮龙","0","0",1,"0",1,"卡牌图片","描述",1,1}
	count := service.CardBaseOfCreate(&card)
	fmt.Println(count)
}

func TestCardBaseOfUpdate(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()

	card := entity.CardBase{}
	card.Image = "卡牌图片"
	count := service.CardBaseOfUpdate(&card)
	fmt.Println(count)
}

func TestCardBaseInfoById(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()

	card := service.CardBaseInfoById(2)
	fmt.Println(card)
}

func TestCardBaseInfoCount(t *testing.T) {
	db := restgo.OpenDBConnect()
	defer db.Close()
	card := entity.CardBase{}
	card.Name = "龙"
	count := service.CardBaseInfoCount(&card)
	fmt.Println(count)
}

func TestCardBaseInfoPage(t *testing.T) {
	/*db := restgo.OpenDBConnect()
	defer db.Close()
	card := entity.CardBase{}
	card.Name = "龙"
	cards := service.CardBaseInfoPage(2,1,&card)
	fmt.Println(cards)*/
	testStr := "c:path/projectName/test"
	subStr := "projectName"
	result := strings.Index(testStr,subStr)+len(subStr)
	fmt.Println(testStr[0:result])
}
