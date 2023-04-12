package transactionEntity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id"`
	Reference   string             `json:"reference"`
	Amount      int                `json:"amount" validate:"required,min=1"`
	CustomerId  primitive.ObjectID `json:"customerId" bson:"customerId" validate:"required"`
	Email       *string            `json:"email" validate:"required"`
	InvoiceId   primitive.ObjectID `json:"invoiceId" bson:"invoiceId" validate:"required"`
	Status      *string            `json:"status" validate:"required"`
	PaymentLink *string            `json:"paymentLink" bson:"paymentLink"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updateAt" bson:"updateAt"`
}

type UpdateTransaction struct {
	Status    *string   `json:"status"`
	UpdatedAt time.Time `json:"updateAt" bson:"updateAt"`
}

type PayStackReq struct {
	Email  *string `json:"email" validate:"required"`
	Amount int     `json:"amount" validate:"required,min=1"`
}

type PayStackRes struct {
	Status  bool `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Authorization string `json:"authorization_url"`
		AccessCode    string `json:"access_code"`
		Reference     string `json:"reference"`
	} `json:"data"`
}
