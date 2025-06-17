package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product 產品模型

// @Description 產品信息
type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty" example:"507f1f77bcf86cd799439011"`
	Name        string             `json:"name" bson:"name" example:"測試產品"`
	Description string             `json:"description" bson:"description" example:"這是一個測試產品的描述"`
	Price       float64            `json:"price" bson:"price" example:"99.99"`
	Quantity    int                `json:"quantity" bson:"quantity" example:"100"`
}
