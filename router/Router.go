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

func SkillRouter(r *gin.Engine){
	r.GET("/skill/findInfoListPage", controller.SkillInfoListPage)
	r.GET("/skill/finAllInfo", controller.SkillAllInfo)
	r.GET("/skill/findInfoById", controller.SkillInfoById)
	r.GET("/skill/delete",controller.SkillOfDelete)
	r.POST("/skill/save",controller.SkillOfSave)
}

func ProfessionRouter(r *gin.Engine){
	r.GET("/profession/findInfoListPage", controller.ProfessionInfoListPage)
	r.GET("/profession/finAllInfo", controller.ProfessionAllInfo)
	r.GET("/profession/findInfoById", controller.ProfessionInfoById)
	r.GET("/profession/delete",controller.ProfessionOfDelete)
	r.POST("/profession/save",controller.ProfessionOfSave)
}

func CardTypeRouter(r *gin.Engine){
	r.GET("/cardType/findInfoListPage", controller.CardTypeInfoListPage)
	r.GET("/cardType/finAllInfo", controller.CardTypeAllInfo)
	r.GET("/cardType/findInfoById", controller.CardTypeInfoById)
	r.GET("/cardType/delete",controller.CardTypeOfDelete)
	r.POST("/cardType/save",controller.CardTypeOfSave)
}

func CardPackageRouter(r *gin.Engine){
	r.GET("/cardPackage/findInfoListPage", controller.CardPackageInfoListPage)
	r.GET("/cardPackage/finAllInfo", controller.CardPackageAllInfo)
	r.GET("/cardPackage/findInfoById", controller.CardPackageInfoById)
	r.GET("/cardPackage/delete",controller.CardPackageOfDelete)
	r.POST("/cardPackage/save",controller.CardPackageOfSave)
}