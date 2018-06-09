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
	page := model.Page{}
	c.Bind(&page)
	c.Bind(&hero)
	fmt.Println("hero:",hero,"page:",page)
	heroList := service.HeroInfoPage(page.PageSize,page.PageNum,&hero)
	c.JSON(http.StatusOK,heroList)
}