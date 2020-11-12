package controller

import (
	"manufacture_supplier_go/middleware"
	"manufacture_supplier_go/server"
	"net/http"
)

// GroovingCuttingInternalBodyQueryByOrderNumber 查询
func GroovingCuttingInternalBodyQueryByOrderNumber(ctx *middleware.Context) {
	orderNumber := ctx.Query("orderNumber")
	userID, ok := getUserID(ctx)
	if !ok {
		return
	}
	rows, err := server.GroovingCuttingInternalBodyQueryWithOrderNumber(orderNumber, userID)
	if err != nil {
		ctx.Failed(http.StatusBadRequest, "查询失败", rows, err.Error())
		return
	}
	ctx.Success(rows, "查询成功", 0)
	return
}

// GroovingCuttingInternalBodyUpsert 新增或者更新
func GroovingCuttingInternalBodyUpsert(ctx *middleware.Context) {
	body := server.GroovingCuttingInternalBody{}
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
	t, err := server.GroovingCuttingInternalBodyUpsert(body)
	upsertResponse(ctx, t, err)
	return
}
