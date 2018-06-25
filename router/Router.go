package router

import (
	"HearthStoneDatabase/controller"
	"github.com/gin-gonic/gin"
)

func HeroRouter(r *gin.Engine){
	r.GET("/hero/findInfoListPage", controller.HeroInfoListPage)
	r.GET("/hero/findInfoById", controller.HeroInfoById)
	r.DELETE("/hero/delete",controller.HeroOfDelete)
	r.POST("/hero/save",controller.HeroOfSave)
}

func CardRouter(r *gin.Engine){
	r.GET("/card/findInfoListPage", controller.CardInfoListPage)
	r.GET("/card/findInfoById", controller.CardInfoById)
	r.DELETE("/card/delete",controller.CardOfDelete)
	r.POST("/card/save",controller.CardOfSave)
}

func SkillRouter(r *gin.Engine){
	r.GET("/skill/findInfoListPage", controller.SkillInfoListPage)
	r.GET("/skill/finAllInfo", controller.SkillAllInfo)
	r.GET("/skill/findInfoById", controller.SkillInfoById)
	r.DELETE("/skill/delete",controller.SkillOfDelete)
	r.POST("/skill/save",controller.SkillOfSave)
}

func ProfessionRouter(r *gin.Engine){
	r.GET("/profession/findInfoListPage", controller.ProfessionInfoListPage)
	r.GET("/profession/finAllInfo", controller.ProfessionAllInfo)
	r.GET("/profession/findInfoById", controller.ProfessionInfoById)
	r.DELETE("/profession/delete",controller.ProfessionOfDelete)
	r.POST("/profession/save",controller.ProfessionOfSave)
}

func CardTypeRouter(r *gin.Engine){
	r.GET("/cardType/findInfoListPage", controller.CardTypeInfoListPage)
	r.GET("/cardType/finAllInfo", controller.CardTypeAllInfo)
	r.GET("/cardType/findInfoById", controller.CardTypeInfoById)
	r.DELETE("/cardType/delete",controller.CardTypeOfDelete)
	r.POST("/cardType/save",controller.CardTypeOfSave)
}

func CardPackageRouter(r *gin.Engine){
	r.GET("/cardPackage/findInfoListPage", controller.CardPackageInfoListPage)
	r.GET("/cardPackage/finAllInfo", controller.CardPackageAllInfo)
	r.GET("/cardPackage/findInfoById", controller.CardPackageInfoById)
	r.DELETE("/cardPackage/delete",controller.CardPackageOfDelete)
	r.POST("/cardPackage/save",controller.CardPackageOfSave)
}

func AuthoriztionRouter(r *gin.Engine){
	r.POST("/authorization/login",controller.Login)
}