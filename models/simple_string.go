// models/simple_string.go
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SimpleString struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Value string             `json:"value" bson:"value"`
}
