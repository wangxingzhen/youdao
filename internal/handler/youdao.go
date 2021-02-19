package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"youdao/internal/server"
	"youdao/util"
)

func YouDao(ctx *gin.Context) {

	params := &server.FYReq{
		Ip: ctx.ClientIP(), //获取客户端ip,用来当作用户唯一id
	}

	err := ctx.ShouldBindJSON(params)
	if err != nil {
		util.ErrorReturn(ctx, util.LackParameter, err.Error())
		return
	}
	if params.Content == "" {
		util.ReturnError(ctx, util.LackParameter)
		return
	}
	res, err := params.YouDaoServer()
	if err != nil {
		fmt.Printf("%+v", err)
		util.ErrorReturn(ctx, util.LackParameter, err.Error())
		return
	}
	//如果为nil，说明与有道词典交互发生错误
	if res == nil {
		util.ReturnError(ctx, util.QueryFails)
		return
	}
	util.Return(ctx, *res, "查询成功")
	return
}
