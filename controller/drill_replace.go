package controller

import (
	"manufacture_supplier_go/middleware"
	"manufacture_supplier_go/server"
	"net/http"
)

// DrillReplaceQueryWithOrderNumber 查询
func DrillReplaceQueryWithOrderNumber(ctx *middleware.Context) {
	orderNumber := ctx.Query("orderNumber")
	userID, ok := getUserID(ctx)
	if !ok {
		return
	}

	list, err := server.DrillReplaceQueryWithOrderNumber(orderNumber, userID)
	if err != nil {
		ctx.Failed(http.StatusBadRequest, "查询失败", nil, err.Error())
		return
	}
	ctx.Success(list, "查询成功", 0)
	return
}

// DrillReplaceInsert 添加
func DrillReplaceInsert(ctx *middleware.Context) {
	body := server.DrillReplace{}
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
	lastID, err := server.DrillReplaceInsertToDB(body)
	if err != nil {
		ctx.Failed(http.StatusBadRequest, "添加失败", nil, err.Error())
		return
	}
	ctx.Success(lastID, "添加成功", 0)
	return
}
