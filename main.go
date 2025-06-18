//go:generate swag init
package main

import (
	"log"
	"net/http"

	"github.com/Jon-GranDen/crud-api/config"
	"github.com/Jon-GranDen/crud-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           字串 API
// @version         1.0
// @description     最小可行性字串 API，支援 GET/POST 並有 Swagger 文件
// @host            localhost:8080
// @BasePath        /

// GetString godoc
// @Summary      取得最新字串
// @Description  取得資料庫中最新一筆字串
// @Tags         string
// @Success      200  {string}  string  "最新字串"
// @Router       /string [get]
func GetString(c *gin.Context) {
	db, err := config.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "資料庫連線失敗"})
		return
	}
	defer db.Close()

	var s string
	err = db.QueryRow("SELECT value FROM simple_strings ORDER BY id DESC LIMIT 1").Scan(&s)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "查無資料"})
		return
	}
	c.String(http.StatusOK, s)
}

// PostString godoc
// @Summary      儲存字串
// @Description  儲存一個新的字串到資料庫
// @Tags         string
// @Accept       plain
// @Produce      json
// @Param        data  body  string  true  "字串內容"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /string [post]
func PostString(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "讀取失敗"})
		return
	}
	str := string(data)

	db, err := config.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "資料庫連線失敗"})
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO simple_strings (value) VALUES ($1)", str)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "資料庫儲存失敗"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "儲存成功"})
}

func main() {
	// 初始化 Swagger 文檔
	docs.SwaggerInfo.Title = "字串 API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()

	r.GET("/string", GetString)
	r.POST("/string", PostString)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
