package controller

import (
	"github.com/gin-gonic/gin"
	"HearthStoneDatabase/model"
	"HearthStoneDatabase/entity"
	"HearthStoneDatabase/service"
	"net/http"
	"strconv"
)

func SkillInfoListPage(c *gin.Context){
	page := model.Page{10,1}
	skill := entity.Skill{}
	c.Bind(&page)
	c.Bind(&skill)
	skillList := service.SkillsInfoPage(page.PageSize,page.PageNum,&skill)
	total := service.SkillInfoCount(&skill)

	c.JSON(http.StatusOK,gin.H{
		"data" : skillList,
		"sumTotal" : total,
		"status" : 1,
	})
}

func SkillAllInfo(c *gin.Context){
	skill := entity.Skill{}
	skillList := service.SkillsInfoPage(0,0,&skill)
	c.JSON(http.StatusOK,gin.H{
		"data" : skillList,
		"status" : 1,
	})
}

func SkillInfoById(c *gin.Context){
	idStr := c.Query("id")
	id,_ := strconv.Atoi(idStr)
	skill := service.CardBaseInfoById(id)
	c.JSON(http.StatusOK,gin.H{
		"data" : skill,
		"status" : 1,
	})
}

func SkillOfSave(c *gin.Context){
	skill := entity.Skill{}
	c.Bind(&skill)
	var status int64
	if skill.Id==0 {
		status = service.SkillOfCreate(&skill)
	}else{
		status = service.SkillOfUpdate(&skill)
	}

	c.JSON(http.StatusOK,gin.H{
		"id" : skill.Id,
		"status" : status,
	})
}

func SkillOfDelete(c *gin.Context){
	idStr := c.Query("id")
	id,_ := strconv.Atoi(idStr)
	status := service.SkillOfDelete(id)
	c.JSON(http.StatusOK,gin.H{
		"status" : status,
	})
}