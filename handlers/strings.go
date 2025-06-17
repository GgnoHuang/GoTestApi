package handlers

import (
	"context"
	"net/http"

	"github.com/Jon-GranDen/crud-api/config"
	"github.com/Jon-GranDen/crud-api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// POST /api/v1/strings
// @Summary 儲存字串
// @Description 儲存一段使用者提供的字串
// @Tags strings
// @Accept json
// @Produce json
// @Param string body map[string]string true "字串內容（{ \"value\": \"你好\" }）"
// @Success 201 {object} models.SimpleString
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /strings [post]
func PostString(c *gin.Context) {
	var input struct {
		Value string `json:"value" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "請提供 value 字串"})
		return
	}

	str := models.SimpleString{
		ID:    primitive.NewObjectID(),
		Value: input.Value,
	}

	_, err := config.DB.Collection("strings").InsertOne(context.Background(), str)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "資料儲存失敗"})
		return
	}

	c.JSON(http.StatusCreated, str)
}

// GET /api/v1/strings
// @Summary 取得所有字串
// @Description 取得所有已儲存的字串資料
// @Tags strings
// @Produce json
// @Success 200 {array} models.SimpleString
// @Failure 500 {object} map[string]string
// @Router /strings [get]
func GetAllStrings(c *gin.Context) {
	var strings []models.SimpleString

	cursor, err := config.DB.Collection("strings").Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "資料讀取失敗"})
		return
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &strings); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "解析資料錯誤"})
		return
	}

	c.JSON(http.StatusOK, strings)
}
