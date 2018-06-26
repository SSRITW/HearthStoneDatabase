package restgo

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var AccessTokenRedisClient *redis.Pool
var RefreshTokenRedisClient *redis.Pool

func initAccessTokenRedisPool() {
	// 建立连接池
	AccessTokenRedisClient = &redis.Pool{
		MaxIdle : ConfigDataSource.Redis.MaxIdle,
		MaxActive : ConfigDataSource.Redis.MaxActive,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ConfigDataSource.Redis.Host,redis.DialPassword(ConfigDataSource.Redis.Password))
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", ConfigDataSource.Redis.AccessTokenDb)
			return c, nil
		},
	}
}

func initRefreshTokenRedisPool() {
	// 建立连接池
	RefreshTokenRedisClient = &redis.Pool{
		MaxIdle : ConfigDataSource.Redis.MaxIdle,
		MaxActive : ConfigDataSource.Redis.MaxActive,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ConfigDataSource.Redis.Host,redis.DialPassword(ConfigDataSource.Redis.Password))
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", ConfigDataSource.Redis.RefreshTokenDb)
			return c, nil
		},
	}
}
