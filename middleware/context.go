package middleware

import (
	"manufacture_supplier_go/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Context 自定义Context
type Context struct {
	*gin.Context
	logger *logrus.Logger
}

// HandlerFunc ...
type HandlerFunc func(*Context)

// CustomContext 添加自定义字段到context
func CustomContext(handler HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customCtx := &Context{}
		customCtx.Context = ctx
		customCtx.logger = util.NewLog()
		handler(customCtx)
	}
}

// ResData 请求成功
type ResData struct {
	code int
	data interface{}
	msg  string
}

// Success  请求成功
func (c *Context) Success(obj interface{}, msg string, code int) {
	resData := gin.H{
		"code": code,
		"data": obj,
		"msg":  msg,
	}
	c.JSON(http.StatusOK, resData)
}

// Failed  请求成功
func (c *Context) Failed(code int, obj ResData) {
	c.JSON(code, obj)
}
