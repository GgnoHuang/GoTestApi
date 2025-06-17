package handlers

import (
	"context"
	"log"
	"net/http"

	// "crud-api/config"
	// "crud-api/models"

	"github.com/Jon-GranDen/crud-api/config"
	"github.com/Jon-GranDen/crud-api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @Summary 創建新產品
// @Description 創建一個新的產品
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.Product true "產品信息"
// @Success 201 {object} models.Product
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	log.Println("開始創建產品...")

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		log.Printf("JSON 綁定失敗: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("接收到的產品數據: %+v", product)

	result, err := config.DB.Collection("products").InsertOne(context.Background(), product)
	if err != nil {
		log.Printf("數據庫插入失敗: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	product.ID = result.InsertedID.(primitive.ObjectID)
	log.Printf("產品創建成功，ID: %s", product.ID.Hex())

	c.JSON(http.StatusCreated, product)
}

// @Summary 獲取所有產品
// @Description 獲取所有產品列表
// @Tags products
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func GetProducts(c *gin.Context) {
	var products []models.Product
	cursor, err := config.DB.Collection("products").Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &products); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// @Summary 獲取單個產品
// @Description 通過ID獲取產品
// @Tags products
// @Produce json
// @Param id path string true "產品ID"
// @Success 200 {object} models.Product
// @Router /products/{id} [get]
func GetProduct(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var product models.Product
	err = config.DB.Collection("products").FindOne(context.Background(), bson.M{"_id": id}).Decode(&product)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// @Summary 更新產品
// @Description 通過ID更新產品
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "產品ID"
// @Param product body models.Product true "產品信息"
// @Success 200 {object} models.Product
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"name":        product.Name,
			"description": product.Description,
			"price":       product.Price,
			"quantity":    product.Quantity,
		},
	}

	result := config.DB.Collection("products").FindOneAndUpdate(
		context.Background(),
		bson.M{"_id": id},
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)

	if result.Err() != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := result.Decode(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// @Summary 刪除產品
// @Description 通過ID刪除產品
// @Tags products
// @Produce json
// @Param id path string true "產品ID"
// @Success 200 {object} map[string]string
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result, err := config.DB.Collection("products").DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
