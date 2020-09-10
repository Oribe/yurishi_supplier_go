package main

import (
	"fmt"
	"manufacture_supplier_go/controller"
	"manufacture_supplier_go/database"

	"github.com/gin-gonic/gin"
)

func main() {
	defer database.NewDB().Close()
	gin.ForceConsoleColor()
	router := gin.Default()

	tool := router.Group("/tool")
	{
		tool.POST("/login", controller.Login)
	}

	err := router.Run()
	if err != nil {
		fmt.Println("start failed：", err)
	}
}
