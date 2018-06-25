package main

import (
	"github.com/gin-gonic/gin"
	"HearthStoneDatabase/router"
	"HearthStoneDatabase/restgo"
	"github.com/gin-contrib/cors"
)

func main() {
	restgo.ConfigDataSource.InitData() //读取配置文件的数据

	restgo.OpenDBConnect()
	restgo.InitAccessTokenRedisPool()
	restgo.InitRefreshTokenRedisPool()
	restgo.InitCasbin()

	r := gin.Default()
	r.Use(cors.Default())	//增加跨域支持
	r.Use(restgo.NewAuthorizer(restgo.Enforcer))	//增加权限验证

	router.HeroRouter(r)
	router.CardRouter(r)
	router.CardPackageRouter(r)
	router.CardTypeRouter(r)
	router.ProfessionRouter(r)
	router.SkillRouter(r)
	router.AuthoriztionRouter(r)

	r.Run("127.0.0.1:80") // listen and serve on 127.0.0.1:80

}