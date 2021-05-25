package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"
	"youdao/internal/data"
	"youdao/internal/data/test"
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

var num int

func Test(ctx *gin.Context) {
	insert := func(ucId int) {
		db, err := data.GetDB()

		if err != nil {
			fmt.Println("获取数据库句柄失败", fmt.Sprintf("%+v", err))
			return
		}
		da := test.UcInshoot{
			UcID: ucId,
			InID: rand.Uint32(),
		}

		db.Transaction(func(tx *gorm.DB) (err error) {
			test.UcInshootMgr(tx).InsertData(da)
			//time.Sleep(10 * time.Second)
			da.UcID++
			test.UcInshootMgr(tx).InsertData(da)
			//time.Sleep(10 * time.Second)
			da.UcID++
			test.UcInshootMgr(tx).InsertData(da)
			//time.Sleep(10 * time.Second)
			da.UcID++
			test.UcInshootMgr(tx).InsertData(da)
			return
		})
		num += 4
		fmt.Println("已插入" + strconv.Itoa(num) + "----------当前协程数量为：" + strconv.Itoa(runtime.NumGoroutine()))
	}
	var wg sync.WaitGroup
	for i := 0; i < 250000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			insert(i * 10)
		}(i)
	}
	wg.Wait()

	util.Return(ctx, nil, "查询成功")
	return
}

func GetGoroutine(ctx *gin.Context) {
	t := func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(time.Second)
				fmt.Println("=====================================================" + "当前的协程数量为" + strconv.Itoa(runtime.NumGoroutine()) + "====================================")
			}
		}
	}
	con, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	go t(con)

	select {
	case <-con.Done():
		return
	}
}
