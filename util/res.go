package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//对应前端数据交互
//单条数据
func Return(ctx *gin.Context, i interface{}, msg string) {
	if i == false {
		i = make([]string, 0)
	}
	var data = map[string]interface{}{
		"msg":  msg,
		"list": i,
	}
	M := make(map[string]interface{}, 4)
	M = returnsParameter(data, ctx)
	ctx.JSON(http.StatusOK, M)
	return
}

//多条数据
func ReturnList(ctx *gin.Context, i interface{}, pagination interface{}, msg string) {
	var dta = map[string]interface{}{
		"msg":        msg,
		"list":       i,
		"pagination": pagination,
	}
	M := make(map[string]interface{}, 4)
	M = returnsParameter(dta, ctx)
	ctx.JSON(http.StatusOK, M)
	return
}

//单条数据带数据字典
func ReturnMap(ctx *gin.Context, i interface{}, maps interface{}, msg string) {
	var data = map[string]interface{}{
		"list": i,
		"maps": maps,
		"msg":  msg,
	}
	M := make(map[string]interface{}, 4)
	M = returnsParameter(data, ctx)
	ctx.JSON(http.StatusOK, M)
	return
}

//多条数据带数据字典
func ReturnListMap(ctx *gin.Context, i interface{}, maps interface{}, pagination interface{}, msg string) {
	var data = map[string]interface{}{
		"list":       i,
		"maps":       maps,
		"msg":        msg,
		"pagination": pagination,
	}
	M := make(map[string]interface{}, 4)
	M = returnsParameter(data, ctx)
	ctx.JSON(http.StatusOK, M)
	return
}

func ErrorReturn(ctx *gin.Context, code int, msg string) {
	var data = map[string]interface{}{
		"msg": msg,
	}
	M := make(map[string]interface{})
	M = returnsParameter(data, ctx)

	if len(msg) > 0 {
		M["data"] = data
	}
	if code > 0 {
		M["status"] = "fail"
		M["code"] = code
	}
	ctx.JSON(http.StatusOK, M)
	return
}

func ReturnError(ctx *gin.Context, code int) {
	msg := GetMsg(code)
	ErrorReturn(ctx, code, msg)
	return
}

func returnsParameter(data map[string]interface{}, ctx *gin.Context) map[string]interface{} {
	return map[string]interface{}{
		"data":     data,
		"code":     0,
		"status":   "success",
		"taketime": getTakeTime(ctx.GetInt64("program_execution_time")),
	}
}

func getTakeTime(startTime int64) float64 {
	return (float64(time.Now().UnixNano()) - float64(startTime)) / 1e6
}
