package invoiceEntity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	Id         primitive.ObjectID `json:"_id" bson:"_id"`
	CustomerId primitive.ObjectID `json:"customerId" validate:"required"`
	IsPaid     *bool              `json:"isPaid" bson:"isPaid" validate:"required"`
	Item       []SaleItem         `json:"items" validate:"required"`
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time          `json:"updateAt" bson:"updateAt"`
}

type UpdateInvoice struct {
	IsPaid    *bool     `json:"isPaid" bson:"isPaid"`
	UpdatedAt time.Time `json:"updateAt" bson:"updateAt"`
}

type SaleItem struct {
	Name      string             `json:"name" validate:"required"`
	Quantity  int                `json:"quantity"`
	Price     int                `json:"price"`
	ProductId primitive.ObjectID `json:"productId" bson:"productId" validate:"required"`
}
