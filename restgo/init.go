package restgo

func init(){
	ConfigDataSource.initConfig() //读取配置文件的信息
	initDBConnect() //初始化数据库连接
	initAccessTokenRedisPool() //初始化redis的连接池
	initRefreshTokenRedisPool()
	initCasbin() //初始化casbin的连接
}
