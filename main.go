package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"youdao/routers"
	"youdao/util"
)

func main() {
	//加载配置文件
	util.InitConf()

	//初始化gin
	r := gin.Default()
	routers.Init(r)
	host := viper.GetString("web.host")
	port := viper.GetString("web.port")
	if host == "" {
		host = "0.0.0.0"
	}
	if port == "" {
		port = "11111"
	}
	util.EndLessServer(host+":"+port, r)
}
