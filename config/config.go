package config

import (
	"fmt"
	"hh_blog/utils/dir"
	"hh_blog/utils/error_check"
	"os"

	"github.com/spf13/viper"
)

var (
	Env string

	MysqlDefHost string
	MysqlDefPort int
	MysqlDefName string
	MysqlDefUser string
	MysqlDefPwd  string

	RedisDefHost string
	RedisDefPort int

	GinListenPort string
	GinMode       string
)

func ConfInit() {
	GetConf()
	GetMysqlConf()
	GetRedisConf()
	GetBannerConf()
	GetGinConf()
	GetUnixSecConf()
}

func GetConf() {
	env := GetEnvConf()
	switch env {
	case "prod":
		viper.SetConfigName("config")
	case "local":
		viper.SetConfigName("config_local")
	}
	viper.ReadInConfig()
}

func GetEnvConf() string {
	fmt.Println(dir.GetAppPath())
	viper.SetConfigName("env")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./conf")
	viper.ReadInConfig()
	env := viper.GetString("env.mode")
	if env == "" {
		panic("get env mode error")
	}
	return env
}

func GetMysqlConf() {
	// 获得mysql默认数据库
	MysqlDefHost = viper.GetString("connect.mysql.defalut.host")
	MysqlDefPort = viper.GetInt("connect.mysql.defalut.port")
	MysqlDefName = viper.GetString("connect.mysql.defalut.name")
	MysqlDefUser = viper.GetString("connect.mysql.defalut.user")
	MysqlDefPwd = viper.GetString("connect.mysql.defalut.pwd")
}

func GetRedisConf() {
	// 获得redis默认数据库
	RedisDefHost = viper.GetString("connect.redis.defalut.host")
	RedisDefPort = viper.GetInt("connect.redis.defalut.port")
}

func GetBannerConf() {
	file, err := os.ReadFile("./conf/banner.txt")
	error_check.PrintErrorMsg(err, "get banner txt error")
	fmt.Println(string(file))
}

func GetGinConf() {
	GinListenPort = viper.GetString("gin.listen.port")
	GinMode = viper.GetString("gin.mode")
}

var (
	UnixSecDtMax   int
	UnixSecGlueMax int
)

func GetUnixSecConf() {
	UnixSecDtMax = viper.GetInt("unixSec.dt.max")
	UnixSecGlueMax = viper.GetInt("unixSec.glue.max")
}
