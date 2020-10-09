package controller

import (
	"manufacture_supplier_go/constant"
	"manufacture_supplier_go/middleware"
	"net/http"
)

func getUserID(ctx *middleware.Context) (int, bool) {
	id, exist := ctx.Get("userId")
	if !exist {
		ctx.Failed(http.StatusUnauthorized, constant.LoginStatusExpiredMsg, nil)
		return 0, exist
	}
	userID := id.(int)
	return userID, exist
}

func unauthorized() {

}
