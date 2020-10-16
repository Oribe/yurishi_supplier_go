package controller

import (
	"manufacture_supplier_go/middleware"
	"manufacture_supplier_go/server"
	"net/http"
)

// DrillIndexableBodyQueryWidthOrderNumber 查询
func DrillIndexableBodyQueryWidthOrderNumber(ctx *middleware.Context) {
	orderNumber := ctx.Query("orderNumber")

	userID, ok := getUserID(ctx)
	if !ok {
		return
	}

	rows, err := server.DrillIndexableBodyQueryWidthOrderNumber(orderNumber, userID)
	if err != nil {
		ctx.Failed(http.StatusBadRequest, "查询失败", nil, err.Error())
		return
	}

	ctx.Success(rows, "查询成功", 0)
	return
}

// DrillIndexableBodyInsert 新增
func DrillIndexableBodyInsert(ctx *middleware.Context) {
	body := &server.DrillIndexableBody{}
	err := ctx.ShouldBindJSON(body)
	if err != nil {
		ctx.Failed(http.StatusBadRequest, "参数错误", nil, err.Error())
		return
	}
	userID, ok := getUserID(ctx)
	if !ok {
		return
	}
	body.Submitter = userID
	lastInsertID, err := server.DrillIndexableBodyInsert(body)
	if err != nil {
		ctx.Failed(http.StatusBadRequest, "添加失败", nil, err.Error())
		return
	}
	ctx.Success(lastInsertID, "添加成功", 0)
	return
}
