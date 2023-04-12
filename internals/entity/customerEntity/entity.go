package customerEntity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName *string            `json:"firstName" validate:"required,min=3"`
	LastName  *string            `json:"lastName" validate:"required,min=3"`
	Email     string             `json:"email" validate:"required,email"`
	Location  string             `json:"location"`
	StaffId   string             `json:"staffId" bson:"staffId" validate:"required"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updateAt" bson:"updateAt"`
}

type UpdateCustomer struct {
	FirstName *string   `json:"firstName" validate:"min=3"`
	LastName  *string   `json:"lasttName" validate:"min=3"`
	UpdatedAt time.Time `json:"updateAt" bson:"updateAt"`
}
