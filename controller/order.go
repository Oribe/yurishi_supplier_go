package controller

import (
	"manufacture_supplier_go/middleware"
	"manufacture_supplier_go/server"
	"manufacture_supplier_go/util"
	"net/http"
)

// OrderSearch 订单查询
func OrderSearch(ctx *middleware.Context) {
	startTime, ok := ctx.GetQuery("startTime")
	if !ok {
		startTime = util.ParseTimestamp(0)
	}

	endTime, ok := ctx.GetQuery("endTime")
	if !ok {
		endTime = util.NowTime()
	}

	userID, exists := ctx.Get("userId")
	if !exists {
		ctx.Failed(http.StatusBadRequest, util.LoginStatusExpiredMsg, nil)
		return
	}
	uid := userID.(int64)
	orderList, err := server.OrderQuery(startTime, endTime, uid)
	if err != nil {
		ctx.Logger.Errorln(err.Error())
		ctx.Failed(http.StatusBadRequest, "查询失败", err.Error())
	}

	ctx.Success(orderList, "查询成功", 0)
}
