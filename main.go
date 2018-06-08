package main

import (
	"github.com/gin-gonic/gin"
	"HearthStoneDatabase/router"
)

func main() {
	r := gin.Default()
	router.HeroRouter(r)
	r.Run("127.0.0.1:8080") // listen and serve on 0.0.0.0:8080
}