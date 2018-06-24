package restgo

import (
	"github.com/casbin/casbin"
	"os"
	"github.com/casbin/gorm-adapter"
	"path"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Enforcer *casbin.Enforcer

//初始化权限管理
func InitCasbin(){
	wr, _ := os.Getwd()
	adapter := gormadapter.NewAdapter(ConfigDateSource.Database.DriveName, ConfigDateSource.Database.ConnetionURL, true)
	Enforcer = casbin.NewEnforcer(path.Join(wr,"config", "auth_model.conf"),adapter)
	Enforcer.LoadPolicy()
}

func NewAuthorizer(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		//从上一个中间件获得username
		uname := c.Request.FormValue("token")
		//修改 check方法，增加一个用户名参数
		if uname=="" || !checkPermission(c.Request, uname) {
			c.JSON(http.StatusUnauthorized,gin.H {
				"msg" : "没有权限访问",
				"status" : 1,
			})
			c.Abort()
		}
	}
}

func checkPermission(r *http.Request, uname string) bool {
	user := uname
	method := r.Method
	path := r.URL.Path
	return Enforcer.Enforce(user, path, method)
}
