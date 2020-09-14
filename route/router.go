package route

import (
	"manufacture_supplier_go/controller"
	"manufacture_supplier_go/middleware"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// Router 定义路由
func init() {
	router = gin.Default()

	router.Use(middleware.Cors())
	router.Use(middleware.Logger())
	{
		handler := middleware.CustomContext
		tool := router.Group("/tool/interface")
		{
			tool.POST("/login", handler(controller.Login))
		}
	}
}

// NewRouter 构造router
func NewRouter() *gin.Engine {
	return router
}
