package controller

import (
	"manufacture_supplier_go/constant"
	"manufacture_supplier_go/middleware"
	"manufacture_supplier_go/util"
	"net/http"
	"strconv"
)

func getUserID(ctx *middleware.Context) (string, bool) {
	id, exist := ctx.Get("userId")
	if !exist {
		ctx.Failed(http.StatusUnauthorized, constant.LoginStatusExpiredMsg, nil, "")
		return "", exist
	}
	userID := id.(int)
	return strconv.Itoa(userID), exist
}

func upsertResponse(ctx *middleware.Context, t int, err error) {
	if t == util.INSERT {
		if err != nil {
			ctx.Failed(http.StatusBadRequest, "添加失败", nil, err.Error())
			return
		}
		ctx.Success(nil, "添加成功", 0)
		return
	}
	if t == util.UPDATE {
		if err != nil {
			ctx.Failed(http.StatusBadRequest, "添加失败", nil, err.Error())
			return
		}
		ctx.Success(nil, "添加成功", 0)
		return
	}
	ctx.Failed(http.StatusBadRequest, "请求失败", nil, err.Error())
	return
}
