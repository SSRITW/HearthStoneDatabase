package controller

import (
	"github.com/gin-gonic/gin"
	"HearthStoneDatabase/entity"
	"HearthStoneDatabase/service"
	"HearthStoneDatabase/restgo"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"net/http"
)

func Login(c *gin.Context){
	user := entity.User{}
	c.Bind(&user)
	userInfo := service.UserInfoByLoginNameAndPassword(&user)

	redisAccessConn := restgo.AccessTokenRedisClient.Get()
	defer redisAccessConn.Close()

	accessToken := restgo.GetToken()
	if _, err := redisAccessConn.Do("HMSET",redis.Args{}.Add(accessToken).AddFlat(&userInfo)...); err!=nil {
		fmt.Println("HMSET ERROR:",err.Error())

		c.JSON(http.StatusInternalServerError,gin.H{
			"msg" : "fail",
			"status" : 0,
		})

		return
	}

	redisRefreshConn := restgo.RefreshTokenRedisClient.Get()
	defer redisRefreshConn.Close()

	refreshToken := restgo.GetToken()
	if _,err := redisRefreshConn.Do("set",refreshToken,userInfo.Id); err!=nil {
		fmt.Println("HMSET ERROR:",err.Error())

		c.JSON(http.StatusInternalServerError,gin.H{
			"msg" : "fail",
			"status" : 0,
		})

		return
	}

	c.JSON(http.StatusOK,gin.H{
		"refreshToken" : refreshToken,
		"accessToken" : accessToken,
		"status" : 1,
	})
}
