package userEntity

import (
	"mime/multipart"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateUserReq struct {
	Id           primitive.ObjectID `form:"_id" bson:"_id"`
	FirstName    string             `form:"firstName" bson:"firstName" validate:"required,min=3"`
	LastName     string             `form:"lastName" bson:"lastName" validate:"required,min=3"`
	Email        string             `form:"email" validate:"required,email"`
	Phone        *string            `form:"phone"`
	ProfileImage *string            `form:"profileImage" bson:"profileImage"`
	Password     string             `form:"password" validate:"required"`
	RefreshToken string             `form:"refreshToken" bson:"refreshToken"`
	CreatedAt    time.Time          `form:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time          `form:"updateAt" bson:"updateAt"`
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

type ImageFile struct {
	File multipart.File `json:"file,omitempty"`
}

type UpdateUserReq struct {
	FirstName    *string   `json:"firstName" bson:"firstName" validate:"min=3"`
	LastName     *string   `json:"lastName" bson:"lastName" validate:"min=3"`
	ProfileImage *string   `json:"profileImage" bson:"profileImage"`
	RefreshToken *string   `json:"refreshToken" bson:"refreshToken"`
	UpdatedAt    time.Time `json:"updateAt" bson:"updateAt"`
}
