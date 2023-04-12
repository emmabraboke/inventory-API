package userEntity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateUserReq struct {
	Id           primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName    string             `json:"firstName" bson:"firstName" validate:"required,min=3"`
	LastName     string             `json:"lastName" bson:"lastName" validate:"required,min=3"`
	Email        string             `json:"email" validate:"required,email"`
	Phone        *string            `json:"phone"`
	ProfileImage *string            `json:"profileImage" bson:"profileImage"`
	Password     string             `json:"password" validate:"required"`
	RefreshToken string             `json:"refreshToken" bson:"refreshToken"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updateAt" bson:"updateAt"`
}

type CreateUserRes struct {
	Id           primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName    *string            `json:"firstName"`
	LastName     *string            `json:"lastName"`
	Email        string             `json:"email"`
	Phone        *string            `json:"phone"`
	ProfileImage *string            `json:"profileImage"`
	CreatedAt    time.Time          `json:"createdAt"`
	UpdatedAt    time.Time          `json:"updateAt"`
}

type Login struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserReq struct {
	FirstName    *string   `json:"firstName" bson:"firstName" validate:"min=3"`
	LastName     *string   `json:"lastName" bson:"lastName" validate:"min=3"`
	ProfileImage *string   `json:"profileImage" bson:"profileImage"`
	RefreshToken *string   `json:"refreshToken" bson:"refreshToken"`
	UpdatedAt    time.Time `json:"updateAt" bson:"updateAt"`
}
