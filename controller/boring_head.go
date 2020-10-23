package controller

import (
	"manufacture_supplier_go/middleware"
	"manufacture_supplier_go/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BoringHeadQuery 查询
func BoringHeadQuery(ctx *middleware.Context) {
	orderNumber := ctx.Query("orderNumber")
	userID, ok := getUserID(ctx)
	if !ok {
		return
	}
	boringHeadList, err := server.BoringHeadQuery(orderNumber, userID)
	if err != nil {
		ctx.Failed(http.StatusBadRequest, err.Error(), nil, err.Error())
		return
	}
	ctx.Success(boringHeadList, "查询成功", 0)
	return
}

// BoringHeadInsert 新增
func BoringHeadInsert(ctx *middleware.Context) {
	body := server.BoringHead{}
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

	lastInserID, err := server.BoringHeadInsert(&body)
	if err != nil {
		ctx.Failed(http.StatusBadRequest, err.Error(), nil, err.Error())
		return
	}

	resData := gin.H{
		"id": lastInserID,
	}

	ctx.Success(resData, "添加成功", 0)
	return
}
