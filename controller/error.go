package controller

import (
	"manufacture_supplier_go/middleware"
	"net/http"
)

// ErrorRecord 错误收集
func ErrorRecord(ctx *middleware.Context) {
	ctx.Success(http.StatusOK, "", 0)
	return
}
