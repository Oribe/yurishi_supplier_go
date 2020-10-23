package controller

import (
	"manufacture_supplier_go/constant"
	"manufacture_supplier_go/middleware"
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
