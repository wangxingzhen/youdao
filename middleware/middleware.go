package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

//记录接口执行时间
func RecordingTime() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("program_execution_time", time.Now().UnixNano())
		ctx.Next()
	}
}
