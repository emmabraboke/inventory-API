package productEntity

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        *string            `json:"name" validate:"required"`
	Description *string            `json:"description"`
	Price       *int               `json:"price"`
	Quantity    *int               `json:"quantity"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updateAt" bson:"updateAt"`
}

type UpdateProduct struct {
	Name        *string   `json:"name" validate:"required"`
	Description *string   `json:"description"`
	Price       *int      `json:"price"`
	Quantity    *int      `json:"quantity"`
	UpdatedAt   time.Time `json:"updateAt" bson:"updateAt"`
}
