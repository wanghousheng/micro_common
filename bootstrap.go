package common

import (
	"fmt"
	"github.com/wanghousheng/micro_common/cache"
	"github.com/wanghousheng/micro_common/database"
	"github.com/wanghousheng/micro_common/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func SetupDB(config *MysqlConfig) {
	var dbConfig gorm.Dialector
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
		config.Charset,
	)
	dbConfig = mysql.New(mysql.Config{
		DSN: dsn,
	})
	// 连接数据库
	database.Connect(dbConfig)
	// 设置最大连接数
	database.SQLDB.SetMaxOpenConns(config.OpenConnections)
	// 设置最大空闲连接数
	database.SQLDB.SetMaxIdleConns(config.IdleConnections)
	// 设置每个链接的过期时间
	database.SQLDB.SetConnMaxLifetime(time.Duration(config.LifeSeconds) * time.Second)
}

// SetupRedis 初始化 Redis
func SetupRedis(config *RedisConfig) {
	// 建立 Redis 连接
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.Host, config.Port),
		config.User,
		config.Password,
		config.Database,
	)
}

// SetupCache 缓存
func SetupCache(config *CachedConfig) {
	// 初始化缓存专用的redis
	rds := cache.NewRedisStore(
		fmt.Sprintf("%v:%v", config.Config.Host, config.Config.Port),
		config.Config.User,
		config.Config.Password,
		config.Config.Database,
		config.Config.Prefix,
	)
	cache.InitWithCacheStore(rds)
}
