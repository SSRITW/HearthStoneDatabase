package controller

import (
	"github.com/gin-gonic/gin"
	"HearthStoneDatabase/model"
	"HearthStoneDatabase/entity"
	"HearthStoneDatabase/service"
	"net/http"
)

func CardPackageInfoListPage(c *gin.Context){
	page := model.Page{10,1}
	cardPackage := entity.CardPackage{}
	c.Bind(&cardPackage)
	c.Bind(&page)
	list := service.CardPackageInfoPage(page.PageSize,page.PageNum,&cardPackage)
	total := service.CardPackageInfoCount(&cardPackage)

	c.JSON(http.StatusOK,gin.H{
		"data" : list,
		"sumTotal" : total,
		"status" : 1,
	})
}

func CardPackageAllInfo(c *gin.Context){
	cardPackage := entity.CardPackage{}
	list := service.CardPackageInfoPage(0,0,&cardPackage)
	c.JSON(http.StatusOK,gin.H{
		"data" : list,
		"status" : 1,
	})
}

func CardPackageInfoById(c *gin.Context){
	id := c.GetInt("id")
	cardPackage := service.CardPackageInfoById(id)

	c.JSON(http.StatusOK,gin.H{
		"data" : cardPackage,
		"status" : 1,
	})
}

func CardPackageOfSave(c *gin.Context){
	cardPackage := entity.CardPackage{}
	c.Bind(&cardPackage)
	var status int64
	if cardPackage.Id==0 {
		status = service.CardPackageOfCreate(&cardPackage)
	}else{
		status = service.CardPackageOfUpdate(&cardPackage)
	}

	c.JSON(http.StatusOK,gin.H{
		"status" : status,
		"id" : cardPackage.Id,
	})
}

func CardPackageOfDelete(c *gin.Context){
	id := c.GetInt("id")
	status := service.CardPackageOfDelete(id)
	c.JSON(http.StatusOK,gin.H{
		"status" : status,
	})
}