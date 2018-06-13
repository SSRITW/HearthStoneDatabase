package main

import (
	"github.com/gin-gonic/gin"
	"HearthStoneDatabase/router"
	"HearthStoneDatabase/restgo"
	"github.com/gin-contrib/cors"
)

func main() {
	restgo.OpenDBConnect()
	r := gin.Default()
	r.Use(cors.Default())	//增加跨域支持
	router.HeroRouter(r)
	router.CardRouter(r)
	router.CardPackageRouter(r)
	router.CardTypeRouter(r)
	router.ProfessionRouter(r)
	router.SkillRouter(r)
	r.Run("127.0.0.1:80") // listen and serve on 0.0.0.0:8080
}