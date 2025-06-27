//go:generate swag init
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/Jon-GranDen/crud-api/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitDB 初始化資料庫連線
func InitDB() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if password == "" {
		password = "password"
	}
	if dbname == "" {
		dbname = "postgres"
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	return sql.Open("postgres", connStr)
}

// GetString 取得最新字串
func GetString(c *gin.Context) {
	db, err := InitDB()
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

// PostString 儲存新的字串
func PostString(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "讀取失敗"})
		return
	}
	str := string(data)

	db, err := InitDB()
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
	myHost := os.Getenv("MY_HOST")
	if myHost == "" {
		myHost = "localhost:8080"
	}

	// 初始化 Swagger
	docs.SwaggerInfo.Title = "字串 API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = myHost
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "API 服務器正在運行"})
	})
	r.GET("/string", GetString)
	r.POST("/string", PostString)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
