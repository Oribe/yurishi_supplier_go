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
			/**
			 * 用户
			 */
			tool.POST("/login", handler(controller.Login))          // 用户登录
			tool.POST("/logout", handler(controller.Logout))        // 用户登出
			tool.POST("/user/update", handler(controller.UserEdit)) //修改用户信息

			/**
			 * 订单
			 */
			tool.GET("/order/query", handler(controller.OrderSearch)) // 订单查询

			/**
			 * 镗削
			 */
			tool.GET("/boring/head", handler(controller.BoringHeadQuery))   // BoringHead 镗头
			tool.POST("/boring/head", handler(controller.BoringHeadInsert)) // BoringHead 镗头

			/**
			 * 钻削
			 */
			tool.GET("/drill/gun", handler(controller.DrillGunQueryWithOrderNumber)) // 枪钻
			tool.POST("/drill/gun", handler(controller.DrillGunInsert))              // 枪钻

			tool.GET("/drill/indexable/body", handler(controller.DrillIndexableBodyQueryWidthOrderNumber)) // 可转位钻头刀体
			tool.POST("/drill/indexable/body", handler(controller.DrillIndexableBodyInsert))               // 可转位钻头刀体

			tool.GET("/drill/integral", handler(controller.DrillIntegralQueryWithOrderNumber)) // 整体式钻头
			tool.POST("/drill/integral", handler(controller.DrillIntegralInsert))              // 整体式钻头

			tool.GET("/drill/replace", handler(controller.DrillReplaceQueryWithOrderNumber)) // 可更换钻削刀头
			tool.POST("/drill/replace", handler(controller.DrillReplaceInsert))              // 可更换钻削刀头

			tool.GET("/drill/replace/body", handler(controller.DrillReplaceBodyQueryWithOrderNumber)) // 可更换钻削刀头刀体
			tool.POST("/drill/replace/body", handler(controller.DrillIndexableBodyInsert))            // 可更换钻削刀头刀体

			tool.GET("/drill/rotate/blade", handler(controller.DrillRotateBladeQueryWithOrderNumber)) // 可转位钻头刀片
			tool.POST("/drill/rotate/blade", handler(controller.DrillRotateBladeInsert))              // 可转位钻头刀片

			tool.GET("/groovingCutting/blade", handler(controller.GroovingCuttingBladeQueryWithOrderNumber)) // 切槽切断刀片
			tool.GET("/groovingCutting/blade", handler(controller.GroovingCuttingBladeInsert))               // 切槽切断刀片

			/**
			* 错误
			 */
			tool.POST("/error/record", handler(controller.ErrorRecord)) // 错误收集

		}
	}
}

// NewRouter 构造router
func NewRouter() *gin.Engine {
	return router
}
