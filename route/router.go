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
	router.Use(middleware.Auth())
	{
		handler := middleware.CustomContext
		tool := router.Group("/tool/interface")
		{
			tool.POST("/login", handler(controller.Login))          // 用户登录
			tool.POST("/logout", handler(controller.Logout))        // 用户登出
			tool.POST("/user/update", handler(controller.UserEdit)) //修改用户信息
		}
	}
}

// NewRouter 构造router
func NewRouter() *gin.Engine {
	return router
}
