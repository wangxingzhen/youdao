package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"youdao/internal/data"
	"youdao/routers"
	"youdao/util"
)

func main() {
	//配置文件
	util.HookLoad = append(util.HookLoad, util.InitConf)
	//数据库连接
	util.HookLoad = append(util.HookLoad, data.InitData)

	//加载
	util.ReloadHookFunc()

	//初始化gin
	r := gin.Default()

	//性能分析
	pprof.Register(r)

	routers.Init(r)
	host := viper.GetString("web.host")
	port := viper.GetString("web.port")
	if host == "" {
		host = "0.0.0.0"
	}
	if port == "" {
		port = "11111"
	}
	host = "0.0.0.0"
	util.EndLessServer(host+":"+port, r)
}
