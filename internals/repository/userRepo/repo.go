package userRepo

import "inventory/internals/entity/userEntity"

type UserRepository interface {
	CreateUser(req *userEntity.CreateUserReq) error
	GetUsers() ([]*userEntity.CreateUserRes, error)
	GetUserByEmail(email string) (*userEntity.CreateUserReq, error)
	GetUser(id string) (*userEntity.CreateUserRes, error)
	UpdateUser(id string, req *userEntity.UpdateUserReq) (*userEntity.CreateUserRes, error)
	DeleteUser(id string) error
}
