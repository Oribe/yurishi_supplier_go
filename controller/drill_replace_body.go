package controller

import (
	"manufacture_supplier_go/middleware"
	"manufacture_supplier_go/server"
	"net/http"
)

// DrillReplaceBodyQueryWithOrderNumber 查询
func DrillReplaceBodyQueryWithOrderNumber(ctx *middleware.Context) {
	orderNumber := ctx.Query("orderNumber")
	userID, ok := getUserID(ctx)
	if !ok {
		return
	}

	rows, err := server.DrillReplaceBodyQueryWithOrderNumber(orderNumber, userID)
	if err != nil {
		ctx.Failed(http.StatusBadRequest, "查询失败", nil, err.Error())
		return
	}
	ctx.Success(rows, "查询成功", 0)
	return
}

// DrillReplaceBodyInsert 添加
func DrillReplaceBodyInsert(ctx *middleware.Context) {
	body := server.DrillReplaceBody{}
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
	lastID, err := server.DrillReplaceBodyInsert(body)
	if err != nil {
		ctx.Failed(http.StatusBadRequest, "添加失败", nil, err.Error())
		return
	}
	ctx.Success(lastID, "添加成功", 0)
	return
}
