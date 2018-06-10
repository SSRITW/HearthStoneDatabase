package controller

import (
	"github.com/gin-gonic/gin"
	"HearthStoneDatabase/model"
	"HearthStoneDatabase/entity"
	"HearthStoneDatabase/service"
	"net/http"
)

func ProfessionInfoListPage(c *gin.Context){
	page := model.Page{10,1}
	profession := entity.Profession{}
	c.Bind(&profession)
	c.Bind(&page)
	list := service.ProfessionInfoPage(page.PageSize,page.PageNum,&profession)
	total := service.ProfessionInfoCount(&profession)

	c.JSON(http.StatusOK,gin.H{
		"data" : list,
		"sumTotal" : total,
		"status" : 1,
	})
}

func ProfessionAllInfo(c *gin.Context){
	profession := entity.Profession{}
	list := service.ProfessionInfoPage(0,0,&profession)
	c.JSON(http.StatusOK,gin.H{
		"data" : list,
		"status" : 1,
	})
}

func ProfessionInfoById(c *gin.Context){
	id := c.GetInt("id")
	profession := service.ProfessionInfoById(id)

	c.JSON(http.StatusOK,gin.H{
		"data" : profession,
		"status" : 1,
	})
}

func ProfessionOfSave(c *gin.Context){
	profession := entity.Profession{}
	c.Bind(&profession)
	var status int64
	if profession.Id==0 {
		status = service.ProfessionOfCreate(&profession)
	}else{
		status = service.ProfessionOfUpdate(&profession)
	}

	c.JSON(http.StatusOK,gin.H{
		"status" : status,
		"id" : profession.Id,
	})
}

func ProfessionOfDelete(c *gin.Context){
	id := c.GetInt("id")
	status := service.ProfessionOfDelete(id)
	c.JSON(http.StatusOK,gin.H{
		"status" : status,
	})
}