package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // 匿名匯入 PostgreSQL 驅動，觸發初始化
)

// InitDB 初始化資料庫連線
func InitDB() (*sql.DB, error) {
	// 从环境变量获取数据库配置，如果没有则使用默认值
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "dpg-d18k7smmcj7s73a31b20-a.singapore-postgres.render.com"
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "jondb_user"
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "2xyCmsk8tPHtbB4kgEtHkykem4S1g0Uw"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "jondb"
	}

	sslMode := os.Getenv("DB_SSLMODE")
	if sslMode == "" {
		sslMode = "require"
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		dbUser, dbPassword, dbName, dbHost, dbPort, sslMode)

	// 打開資料庫連線（這邊不會馬上連線）
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("無法連線到資料庫: %v", err)
	}

	// 用 Ping 測試連線是否正常（這才是實際連線動作）
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("資料庫連線測試失敗: %v", err)
	}

	fmt.Println("✅ 成功連接到 PostgreSQL!")
	return db, nil
}

func checkErr(err error) {
	if err != nil {
		log.Fatal("錯誤：", err)
	}
}
