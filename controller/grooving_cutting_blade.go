package controller

import (
	"manufacture_supplier_go/middleware"
	"manufacture_supplier_go/server"
	"net/http"
)

// GroovingCuttingBladeQueryWithOrderNumber 查询
func GroovingCuttingBladeQueryWithOrderNumber(ctx *middleware.Context) {
	orderNumber := ctx.Query("orderNumber")
	userID, ok := getUserID(ctx)
	if !ok {
		return
	}
	rows, err := server.GroovingCuttingBladeQueryWithOrderNumber(orderNumber, userID)
	if err != nil {
		ctx.Failed(http.StatusBadRequest, "查询失败", nil, err.Error())
		return
	}
	ctx.Success(rows, "查询成功", 0)
	return
}

// GroovingCuttingBladeInsert 添加
func GroovingCuttingBladeInsert(ctx *middleware.Context) {
	body := server.GroovingCuttingBlade{}
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.Failed(http.StatusBadRequest, "参数错误", nil, err.Error())
		return
	}
	userID, ok := getUserID(ctx)
	if !ok {
		return
	}
	body.Submitter = userID
	id, err := server.GroovingCuttingBladeInsert(body)
	if err != nil {
		ctx.Failed(http.StatusBadRequest, "添加失败", nil, err.Error())
		return
	}
	ctx.Success(id, "添加成功", 0)
	return
}
