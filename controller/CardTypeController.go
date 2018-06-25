package controller

import (
	"github.com/gin-gonic/gin"
	"HearthStoneDatabase/model"
	"HearthStoneDatabase/entity"
	"HearthStoneDatabase/service"
	"net/http"
	"strconv"
)

func CardTypeInfoListPage(c *gin.Context){
	page := model.Page{10,1}
	cardType := entity.CardType{}
	c.Bind(&cardType)
	c.Bind(&page)
	list := service.CardTypeInfoPage(page.PageSize,page.PageNum,&cardType)
	total := service.CardTypeInfoCount(&cardType)

	c.JSON(http.StatusOK,gin.H{
		"data" : list,
		"sumTotal" : total,
		"status" : 1,
	})
}

func CardTypeAllInfo(c *gin.Context){
	cardType := entity.CardType{}
	list := service.CardTypeInfoPage(0,0,&cardType)
	c.JSON(http.StatusOK,gin.H{
		"data" : list,
		"status" : 1,
	})
}

func CardTypeInfoById(c *gin.Context){
	idStr := c.Query("id")
	id,_ := strconv.Atoi(idStr)
	cardType := service.CardTypeInfoById(id)

	c.JSON(http.StatusOK,gin.H{
		"data" : cardType,
		"status" : 1,
	})
}

func CardTypeOfSave(c *gin.Context){
	cardType := entity.CardType{}
	c.Bind(&cardType)
	var status int64
	if cardType.Id==0 {
		status = service.CardTypeOfCreate(&cardType)
	}else{
		status = service.CardTypeOfUpdate(&cardType)
	}

	c.JSON(http.StatusOK,gin.H{
		"status" : status,
		"id" : cardType.Id,
	})
}

func CardTypeOfDelete(c *gin.Context){
	idStr := c.Query("id")
	id,_ := strconv.Atoi(idStr)
	status := service.CardTypeOfDelete(id)
	c.JSON(http.StatusOK,gin.H{
		"status" : status,
	})
}