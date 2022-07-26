package main

import (
	"github.com/asim/go-micro/plugins/config/source/consul/v3"
	"github.com/asim/go-micro/v3/config"
	"strconv"
)

// GetConsulConfig 设置配置中心
func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		consul.WithPrefix(prefix),
		consul.StripPrefix(true),
	)
	//配置初始化
	conf, err := config.NewConfig()
	if err != nil {
		return conf, err
	}
	//加载配置
	err = conf.Load(consulSource)
	return conf, err
}

// MysqlConfig mysql 配置信息
type MysqlConfig struct {
	Host            string `json:"host"`
	User            string `json:"user"`
	Password        string `json:"password"`
	Database        string `json:"database"`
	Port            string `json:"port"`
	Charset         string `json:"charset"`
	OpenConnections string `json:"open_connections"` //最大连接数
	IdleConnections string `json:"idle_connections"` //最大空闲连接数
	LifeSeconds     string `json:"life_Seconds"`     //连接过期时间
}

// RedisConfig redis 配置信息
type RedisConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Port     string `json:"port"`
	Prefix   string `json:"prefix"`
}

// GetMysqlFromConsul 获取mysql的配置
func GetMysqlFromConsul(config config.Config, path ...string) *MysqlConfig {
	mysqlConfig := &MysqlConfig{}
	//获取配置
	_ = config.Get(path...).Scan(mysqlConfig)
	return mysqlConfig
}

// GetRedisFromConsul  获取redis的配置
func GetRedisFromConsul(config config.Config, path ...string) *RedisConfig {
	redisConfig := &RedisConfig{}
	//获取配置
	_ = config.Get(path...).Scan(redisConfig)
	return redisConfig
}
