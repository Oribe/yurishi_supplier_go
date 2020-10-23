package controller

import (
	"manufacture_supplier_go/middleware"
	"manufacture_supplier_go/server"
	"net/http"
)

// DrillIntegralQueryWithOrderNumber 查询
func DrillIntegralQueryWithOrderNumber(ctx *middleware.Context) {
	orderNumber := ctx.Query("orderNumber")
	userID, ok := getUserID(ctx)
	if !ok {
		return
	}
	list, err := server.DrillIntegralQueryWithOrderNumber(orderNumber, userID)
	if err != nil {
		ctx.Failed(http.StatusBadRequest, "查询错误", nil, err.Error())
		return
	}
	ctx.Success(list, "查询成功", 0)
	return
}

// DrillIntegralInsert 新增
func DrillIntegralInsert(ctx *middleware.Context) {
	body := server.DrillIntegral{}
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.Failed(http.StatusBadRequest, "参数错误", nil, err.Error())
		return
	}
	id, err := server.DrillIntegralInsertToDB(body)
	if err != nil {
		ctx.Failed(http.StatusBadRequest, "添加失败", nil, err.Error())
		return
	}
	ctx.Success(id, "添加成功", 0)
	return
}
