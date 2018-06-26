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

	result := gin.H{}
	httpStatus := 200

	if userInfo.Id==0 {
		result["msg"] = "用户名不存在或密码错误"
		result["status"] = restgo.LOGIN_FAIL_STATUS
	}else{
		redisAccessConn := restgo.AccessTokenRedisClient.Get()
		defer redisAccessConn.Close()

		accessToken := restgo.GetToken()
		if _, err := redisAccessConn.Do("HMSET",redis.Args{}.Add(accessToken).AddFlat(&userInfo)...); err!=nil {
			fmt.Println("redis HMSET ERROR:",err.Error())

			httpStatus = http.StatusInternalServerError
			result["msg"] = "fail"
			result["status"] = restgo.FAIL
		} else {
			redisRefreshConn := restgo.RefreshTokenRedisClient.Get()
			defer redisRefreshConn.Close()

			refreshToken := restgo.GetToken()
			if _,err := redisRefreshConn.Do("SET",refreshToken,userInfo.Id); err!=nil {
				fmt.Println("redis SET ERROR:",err.Error())

				httpStatus = http.StatusInternalServerError
				result["msg"] = "fail"
				result["status"] = restgo.FAIL
			} else {
				result["refreshToken"] = refreshToken
				result["accessToken"] = accessToken
				result["status"] = restgo.SUCCESS_STATUS
			}
		}
	}

	c.JSON(httpStatus,result)
}
