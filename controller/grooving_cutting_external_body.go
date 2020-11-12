package controller

import (
	"manufacture_supplier_go/middleware"
	"manufacture_supplier_go/server"
	"net/http"
)

// GroovingCuttingExternalBodyInsertToDB 新增
func GroovingCuttingExternalBodyInsertToDB(ctx *middleware.Context) {
	body := server.GroovingCuttingExternalBody{}
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
	id, err := server.GroovingCuttingExternalBodyInsertToDB(body)
	if err != nil {
		ctx.Failed(http.StatusBadRequest, "添加失败", nil, err.Error())
		return
	}
	ctx.Success(id, "添加成功", 0)
	return
}

// GroovingCuttingExternalBodyQueryWithOrderNumber 查询
func GroovingCuttingExternalBodyQueryWithOrderNumber(ctx *middleware.Context) {
	orderNumber := ctx.Query("orderNumber")
	userID, ok := getUserID(ctx)
	if !ok {
		return
	}
	rows, err := server.GroovingCuttingExternalBodyQueryWithOrderNumber(orderNumber, userID)
	if err != nil {
		ctx.Failed(http.StatusBadRequest, "查询失败", err.Error(), err.Error())
		return
	}
	ctx.Success(rows, "查询成功", 0)
	return
}
