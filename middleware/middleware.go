package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"youdao/util"
)

//捕获panic
func TryCatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				util.Fatal("panic: %+v", err)
			}
		}()
		ctx.Next()
	}
}

//记录接口执行时间
func RecordingTime() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("program_execution_time", time.Now().UnixNano())
		ctx.Next()
	}
}
