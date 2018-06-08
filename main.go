package main

import (
	"github.com/gin-gonic/gin"
	"HearthStoneDatabase/entity"
	"fmt"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		hero := entity.Hero{}
		c.Bind(&hero)
		fmt.Println(hero)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("127.0.0.1:8080") // listen and serve on 0.0.0.0:8080
}