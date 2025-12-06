package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	// MongoDB ObjectID
	//  _id as the primary key
	// json:"_id,omitempty" → when encoding to JSON, use _id
	//bson:"_id,omitempty" → when saving to MongoDB, use _id
	//omitempty → skip the field if it’s empty
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty"`
	Quantity    int                `json:"quantity,omitempty"`
	Description string             `json:"description,omitempty"`
	Price       float64            `json:"price,omitempty"`
}
