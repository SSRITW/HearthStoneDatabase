package router

import (
	"HearthStoneDatabase/controller"
	"github.com/gin-gonic/gin"
)

func HeroRouter(r *gin.Engine){
	r.GET("/FindHeroInfoList", controller.HeroInfoListPage)
}