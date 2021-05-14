package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"youdao/internal/handler"
	"youdao/middleware"
)

func Init(r *gin.Engine) {
	//记录接口执行时间
	r.Use(middleware.TryCatch())
	r.Use(middleware.RecordingTime())
	//心跳
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": 0, "message": "success", "data": "pong pong pong wxz !!!"})
	})

	yd := r.Group("fanyi")
	{
		yd.POST("/youdao", handler.YouDao)
	}

	r.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"), "src/youdao/views/*"))
	r.GET("/local", func(c *gin.Context) {
		c.HTML(http.StatusOK, "local.html", gin.H{})
	})
}
