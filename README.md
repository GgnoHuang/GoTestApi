# Go CRUD API with MongoDB and Swagger

這是一個使用 Go 語言開發的 RESTful API 示例，使用 MongoDB 作為數據庫，並集成了 Swagger 文檔。

## 功能特點

- 完整的 CRUD 操作
- MongoDB 數據庫集成
- Swagger API 文檔
- RESTful API 設計
- Gin Web 框架

## 前置要求

- Go 1.21 或更高版本
- MongoDB 運行在 localhost:27017
- Git

## 安裝步驟

1. 克隆項目：
```bash
git clone <your-repository-url>
cd crud-api
```

2. 安裝依賴：
```bash
go mod download
```

3. 生成 Swagger 文檔：
```bash
swag init
```

4. 運行項目：
```bash
go run main.go
```

## API 端點

- POST /api/v1/products - 創建新產品
- GET /api/v1/products - 獲取所有產品
- GET /api/v1/products/:id - 獲取單個產品
- PUT /api/v1/products/:id - 更新產品
- DELETE /api/v1/products/:id - 刪除產品

## Swagger 文檔

訪問 Swagger UI：http://localhost:8080/swagger/index.html

## 數據庫配置

默認配置：
- 數據庫 URL：mongodb://localhost:27017
- 數據庫名稱：product_db
- 集合名稱：products 