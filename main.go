package main

import (
	"github.com/gin-gonic/gin"
	"HearthStoneDatabase/router"
	"HearthStoneDatabase/restgo"
)

func main() {
	restgo.OpenDBConnect()
	r := gin.Default()
	router.HeroRouter(r)
	router.CardRouter(r)
	r.Run("127.0.0.1:80") // listen and serve on 0.0.0.0:8080
}