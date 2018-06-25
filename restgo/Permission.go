package restgo

import (
	"github.com/casbin/casbin"
	"os"
	"github.com/casbin/gorm-adapter"
	"path"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var Enforcer *casbin.Enforcer

//初始化权限管理
func InitCasbin(){
	wr, _ := os.Getwd()
	adapter := gormadapter.NewAdapter(ConfigDataSource.Database.DriveName, ConfigDataSource.Database.ConnetionURL, true)
	Enforcer = casbin.NewEnforcer(path.Join(wr,"config", "auth_model.conf"),adapter)
	Enforcer.LoadPolicy()
}

func NewAuthorizer(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		isPass := false
		accessToken := c.Request.FormValue("access_token")

		if accessToken != "" {
			redisConn := AccessTokenRedisClient.Get()
			defer redisConn.Close() //关闭redis连接
			roleId,err := redis.String(redisConn.Do("HGET",accessToken,"role_id"))
			if err!=nil {
				fmt.Println("redis get roleId error",err.Error())
			}else if checkPermission(c.Request, roleId) { //检查该角色是否拥有权限访问
				isPass = true
			}
		}else if checkPermission(c.Request,""){	//accessToken为空时，即校验他是否访问的是游客可访问的路径
			isPass = true
		}

		if !isPass {
			c.JSON(http.StatusUnauthorized,gin.H {
				"msg" : "没有权限访问",
				"status" : 1,
			})
			c.Abort()
		}
	}
}

func checkPermission(r *http.Request, roleId string) bool {
	method := r.Method
	path := r.URL.Path
	if roleId == "" {
		roleId = VISITOR
	}
	return Enforcer.Enforce(roleId, path, method)
}
