package saleEntity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sale struct {
	Id         primitive.ObjectID `json:"_id" bson:"_id"`
	Name       string             `json:"name" validate:"required"`
	Quantity   int                `json:"quantity"`
	Price      int                `json:"price"`
	CustomerId primitive.ObjectID `json:"customerId" bson:"customerId" validate:"required"`
	InvoiceId  primitive.ObjectID `json:"invoiceId" bson:"invoiceId" validate:"required"`
	ProductId  primitive.ObjectID `json:"productId" bson:"productId" validate:"required"`
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time          `json:"updatedAt" bson:"updatedAt"`
}


