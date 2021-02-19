package util

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"syscall"
	"time"
)

//配置文件的目录
const _confPath = "./conf"

//初始化配置文件
func InitConf() {
	//设置配置文件类型
	viper.SetConfigType("toml")
	//配置文件路径
	viper.AddConfigPath(_confPath)

	files, err := ioutil.ReadDir(_confPath)
	if err != nil {
		Fatal("读取环境配置错误 err(%s)", err.Error())
		panic("初始化配置错误")
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		viper.SetConfigName(f.Name())
		if err = viper.MergeInConfig(); err != nil {
			Fatal("读取环境配置错误 err(%s) filename(%s)", err.Error(), f.Name())
			panic("读取配置出错，请检查后重试")
		}
	}
	return
}

//添加平滑重启，监听端口
func EndLessServer(addr string, r *gin.Engine) {

	endless.DefaultReadTimeOut = time.Second * 20
	endless.DefaultWriteTimeOut = time.Second * 20
	endless.DefaultMaxHeaderBytes = 1 << 20

	srv := endless.NewServer(addr, r)   //新建一个http服务，传入Addr 与 r
	srv.IdleTimeout = time.Second * 10  //空闲超时时间
	srv.ReadTimeout = time.Second * 75  //读取超时时间
	srv.WriteTimeout = time.Second * 75 //写入超时时间
	srv.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
