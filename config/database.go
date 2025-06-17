package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// ConnectDB 連接到 MongoDB
func ConnectDB() {
	log.Println("=== MongoDB 連接檢查開始 ===")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 連接到 MongoDB
	log.Printf("正在嘗試連接到 MongoDB: mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("MongoDB 連接失敗:", err)
	}

	// 檢查連接
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB Ping 失敗:", err)
	}

	DB = client.Database("product_db")
	log.Printf("已選擇數據庫: product_db")

	// 列出所有數據庫
	dbs, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Printf("警告：無法列出數據庫列表: %v", err)
	} else {
		log.Printf("可用的數據庫列表: %v", dbs)
	}

	// 檢查集合
	collections, err := DB.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		log.Printf("警告：無法列出集合: %v", err)
	} else {
		log.Printf("product_db 中的集合: %v", collections)
	}

	// 嘗試創建一個測試文檔
	_, err = DB.Collection("products").InsertOne(ctx, bson.M{"_test": "connectivity_test"})
	if err != nil {
		log.Printf("警告：無法創建測試文檔: %v", err)
	} else {
		// 清理測試文檔
		_, err = DB.Collection("products").DeleteOne(ctx, bson.M{"_test": "connectivity_test"})
		if err != nil {
			log.Printf("警告：無法刪除測試文檔: %v", err)
		}
		log.Println("數據庫寫入測試成功")
	}

	log.Println("=== MongoDB 連接檢查完成 ===")
	log.Println("MongoDB 連接成功並可以正常操作!")
}
