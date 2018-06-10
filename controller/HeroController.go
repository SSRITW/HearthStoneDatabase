package controller

import (
	"github.com/gin-gonic/gin"
	"HearthStoneDatabase/entity"
	"HearthStoneDatabase/model"
	"HearthStoneDatabase/service"
	"net/http"
	"fmt"
)

func HeroInfoListPage(c *gin.Context){
	hero := entity.Hero{}
	page := model.Page{10,1}
	c.Bind(&page)
	c.Bind(&hero)
	heroList := service.HeroInfoPage(page.PageSize,page.PageNum,&hero)
	total := service.HeroInfoCount(&hero)

	c.JSON(http.StatusOK,gin.H{
		"data" : heroList,
		"sumTotal" : total,
		"status" : 1,
	})
}

func HeroInfoById(c *gin.Context){
	id := c.GetInt("id")
	hero := service.HeroInfoById(id)
	c.JSON(http.StatusOK,gin.H{
		"data" : hero,
		"status" : 1,
	})
}

func HeroOfSave(c *gin.Context){
	hero := entity.Hero{}
	c.Bind(&hero)
	var status int64
	if hero.Id==0 {
		status = service.HeroOfCreate(&hero)
	}else{
		status = service.HeroOfUpdate(&hero)
	}

	fmt.Println(hero)
	c.JSON(http.StatusOK,gin.H{
		"status" : status,
		"id" : hero.Id,
	})
}

func HeroOfDelete(c *gin.Context){
	id := c.GetInt("id")
	status := service.HeroOfDelete(id)

	c.JSON(http.StatusOK,gin.H{
		"status" : status,
	})
}