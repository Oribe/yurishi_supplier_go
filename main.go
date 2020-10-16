package main

import (
	"fmt"
	"manufacture_supplier_go/config"
	"manufacture_supplier_go/database"
	"manufacture_supplier_go/route"
	"manufacture_supplier_go/util"
)

func main() {
	// 关闭数据库连接
	defer database.NewDB().Close()
	// 关闭日志文件
	defer util.LogFileClose()

	// 设置配置
	config.Set()

	router := route.NewRouter()
	// err := router.Run(":8081")
	err := router.Run(":7002")
	if err != nil {
		fmt.Println("start failed：", err)
	}
}
