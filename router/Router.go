package router

import (
	"HearthStoneDatabase/controller"
	"github.com/gin-gonic/gin"
)

func HeroRouter(r *gin.Engine){
	r.GET("/hero/findInfoListPage", controller.HeroInfoListPage)
	r.GET("/hero/findInfoById", controller.HeroInfoById)
	r.GET("/hero/delete",controller.HeroOfDelete)
	r.POST("/hero/save",controller.HeroOfSave)
}

func CardRouter(r *gin.Engine){
	r.GET("/card/findInfoListPage", controller.CardInfoListPage)
	r.GET("/card/findInfoById", controller.CardInfoById)
	r.GET("/card/delete",controller.CardOfDelete)
	r.POST("/card/save",controller.CardOfSave)
}