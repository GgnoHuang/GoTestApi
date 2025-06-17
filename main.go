package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Jon-GranDen/crud-api/config"
	"github.com/Jon-GranDen/crud-api/docs"
	"github.com/Jon-GranDen/crud-api/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Product API
// @version         1.0
// @description     This is a sample product CRUD API with MongoDB
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	fmt.Println("Hello, World!")

	// 初始化 Swagger 文檔
	docs.SwaggerInfo.Title = "Product API"
	docs.SwaggerInfo.Description = "This is a sample product CRUD API with MongoDB"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// 連接數據庫
	config.ConnectDB()

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
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		// 產品路由
		products := v1.Group("/products")
		{
			products.POST("", handlers.CreateProduct)
			products.GET("", handlers.GetProducts)
			products.GET("/:id", handlers.GetProduct)
			products.PUT("/:id", handlers.UpdateProduct)
			products.DELETE("/:id", handlers.DeleteProduct)
		}

		v1.POST("/strings", handlers.PostString)
		v1.GET("/strings", handlers.GetAllStrings)

		// 測試路由
		v1.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "API v1 測試端點正常",
			})
		})
	}

	// Swagger 文檔路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 打印所有註冊的路由
	log.Println("=== 註冊的路由 ===")
	log.Printf("GET    /                    (測試首頁)")
	log.Printf("GET    /api/v1/test         (API 測試端點)")
	log.Printf("POST   /api/v1/products     (創建產品)")
	log.Printf("GET    /api/v1/products     (獲取所有產品)")
	log.Printf("GET    /api/v1/products/:id (獲取單個產品)")
	log.Printf("PUT    /api/v1/products/:id (更新產品)")
	log.Printf("DELETE /api/v1/products/:id (刪除產品)")
	log.Printf("GET    /swagger/*any        (Swagger 文檔)")
	log.Println("==================")

	log.Printf("服務器啟動在 http://localhost:8080")
	log.Printf("可以訪問以下 URL 測試服務器：")
	log.Printf("1. 首頁測試：http://localhost:8080/")
	log.Printf("2. API 測試：http://localhost:8080/api/v1/test")
	log.Printf("3. Swagger UI：http://localhost:8080/swagger/index.html")
	log.Printf("4. 產品 API：http://localhost:8080/api/v1/products")

	// 啟動服務器
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
