package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
	"syscall"
	"time"
	"youdao/routers"
)

func main() {
	r := gin.Default()
	routers.Init(r)
	EndLessServer("127.0.0.1:11111", r)
}

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
