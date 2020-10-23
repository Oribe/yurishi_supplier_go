package controller

import (
	"manufacture_supplier_go/middleware"
	"manufacture_supplier_go/server"
	"net/http"
)

// DrillRotateBladeQueryWithOrderNumber 查询
func DrillRotateBladeQueryWithOrderNumber(ctx *middleware.Context) {
	orderNumber := ctx.Query("orderNumber")
	userID, ok := getUserID(ctx)
	if !ok {
		return
	}
	rows, err := server.DrillRotateBladeQueryWithOrderNumber(orderNumber, userID)
	if err != nil {
		ctx.Failed(http.StatusBadRequest, "查询失败", nil, err.Error())
		return
	}
	ctx.Success(rows, "查询成功", 0)
	return
}

// DrillRotateBladeInsert 添加
func DrillRotateBladeInsert(ctx *middleware.Context) {
	body := server.DrillRotateBlade{}
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
	id, err := server.DrillRotateBladeInsert(body)
	if err != nil {
		ctx.Failed(http.StatusBadRequest, "添加失败", nil, err.Error())
		return
	}
	ctx.Success(id, "添加成功", 0)
	return
}
