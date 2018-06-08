package controller

import (
	"github.com/gin-gonic/gin"
	"HearthStoneDatabase/entity"
	"HearthStoneDatabase/model"
	"HearthStoneDatabase/service"
	"net/http"
)

func HeroInfoListPage(c *gin.Context){
	hero := entity.Hero{}
	page := model.Page{}
	c.Bind(&page)
	c.Bind(&hero)
	heroList := service.HeroInfoPage(page.PageSize,page.PageNum,&hero)
	c.JSON(http.StatusOK,heroList)
}