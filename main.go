package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Jon-GranDen/crud-api/config"
	"github.com/Jon-GranDen/crud-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Product API
// @version         1.0
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	fmt.Println("Hello, World!")

	// 初始化 Swagger 文檔
	docs.SwaggerInfo.Title = "Product API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// 連接數據庫
	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("✅ 成功連接到 PostgreSQL!")

	// 創建 Gin 路由
	r := gin.Default()

	// 添加日誌中間件
	r.Use(gin.Logger())

	// 添加一個測試路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API 服務器正在運行",
		})
	})

	// API 路由
	// api := r.Group("/api")
	// v1 := api.Group("/v1")
	// {
	// 產品路由
	// products := v1.Group("/products")
	// {
	// 	products.POST("", handlers.CreateProduct)

	// }

	// v1.POST("/strings", handlers.PostString)
	// v1.GET("/strings", handlers.GetAllStrings)

	// // 測試路由
	// v1.GET("/test", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "API v1 測試端點正常",
	// 	})
	// })
	// }

	// Swagger 文檔路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 打印所有註冊的路由
	log.Println("=== 註冊的路由 ===")

	// 啟動服務器
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
