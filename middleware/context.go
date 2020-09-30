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
	Logger *logrus.Logger
}

// HandlerFunc ...
type HandlerFunc func(*Context)

// CustomContext 添加自定义字段到context
func CustomContext(handler HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customCtx := &Context{}
		customCtx.Context = ctx
		customCtx.Logger = util.NewLog()
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
	if obj == nil {
		obj = gin.H{}
	}
	resData := gin.H{
		"code": code,
		"data": obj,
		"msg":  msg,
	}
	c.JSON(http.StatusOK, resData)
}

// Failed  请求失败
func (c *Context) Failed(code int, msg string, obj interface{}) {
	if obj == nil {
		obj = gin.H{}
	}
	ResData := gin.H{
		"msg":  msg,
		"code": code,
		"data": obj,
	}
	c.JSON(code, ResData)
}
