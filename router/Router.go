package router

import (
	"HearthStoneDatabase/controller"
	"github.com/gin-gonic/gin"
)

func HeroRouter(r *gin.Engine){
	r.GET("/hero/findInfoListPage", controller.FindHeroInfoListPage)
	r.GET("/hero/findInfoById", controller.FindHeroInfoById)
	r.POST("/hero/save",controller.SaveHero)
}

func CardRouter(r *gin.Engine){
	r.GET("/card/findInfoListPage", controller.FindHeroInfoListPage)
	r.GET("/card/findInfoById", controller.FindHeroInfoById)
	r.POST("/card/save",controller.SaveHero)
}