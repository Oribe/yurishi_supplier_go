package controller

import (
	"manufacture_supplier_go/middleware"
	"manufacture_supplier_go/server"
	"net/http"
)

// DrillGunQueryWithOrderNumber 查询
func DrillGunQueryWithOrderNumber(ctx *middleware.Context) {
	orderNumber := ctx.Query("orderNumber")

	userID, ok := getUserID(ctx)
	if !ok {
		return
	}

	data, err := server.DrillGunQueryWithOrder(orderNumber, userID)
	if err != nil {
		ctx.Failed(http.StatusBadRequest, "查询失败", nil, err.Error())
		return
	}

	ctx.Success(data, "查询成功", 0)
	return
}

// DrillGunInsert 新增
func DrillGunInsert(ctx *middleware.Context) {
	body := server.DrillGun{}
	err := ctx.ShouldBindJSON(body)
	if err != nil {
		ctx.Failed(http.StatusBadRequest, "参数错误", nil, err.Error())
		return
	}

	userID, ok := getUserID(ctx)
	if !ok {
		return
	}

	id, err := server.DrillGunInsert(&body, userID)
	if err != nil {
		ctx.Failed(http.StatusBadRequest, "添加失败", nil, err.Error())
		return
	}
	ctx.Success(id, "添加成功", 0)
	return
}
