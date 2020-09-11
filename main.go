package main

import (
	"fmt"
	"manufacture_supplier_go/config"
	"manufacture_supplier_go/database"
	"manufacture_supplier_go/route"
)

func main() {
	// 关闭数据库连接
	defer database.NewDB().Close()

	// 设置配置
	config.Set()

	router := route.NewRouter()
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("start failed：", err)
	}
}
