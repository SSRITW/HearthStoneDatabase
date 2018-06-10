package controller

import (
	"github.com/gin-gonic/gin"
	"HearthStoneDatabase/model"
	"HearthStoneDatabase/service"
	"HearthStoneDatabase/entity"
	"net/http"
)

func CardInfoListPage(c *gin.Context){
	page := model.Page{10,1}
	card := entity.CardBase{}
	c.Bind(&page)
	c.Bind(&card)
	cards := service.CardBaseInfoPage(page.PageSize,page.PageNum,&card)
	total := service.CardBaseInfoCount(&card)

	c.JSON(http.StatusOK,gin.H{
		"data" : cards,
		"sumTotal" : total,
		"status" : 1,
	})
}

func CardInfoById(c *gin.Context){
	id := c.GetInt("id")
	card := service.CardBaseInfoById(id)
	c.JSON(http.StatusOK,gin.H{
		"data" : card,
		"status" : 1,
	})
}

func CardOfSave(c *gin.Context){
	card := entity.CardBase{}
	c.Bind(&card)
	var status int64
	if card.Id==0 {
		status = service.CardBaseOfCreate(&card)
	}else{
		status = service.CardBaseOfUpdate(&card)
	}

	c.JSON(http.StatusOK,gin.H{
		"id" : card.Id,
		"status" : status,
	})
}

func CardOfDelete(c *gin.Context){
	id := c.GetInt("id")
	status := service.CardBaseOfDelete(id)

	c.JSON(http.StatusOK,gin.H{
		"status" : status,
	})
}
