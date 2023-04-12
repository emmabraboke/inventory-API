package userService

import (
	// "fmt"
	"encoding/json"
	"fmt"
	"inventory/internals/entity/userEntity"
	"inventory/internals/repository/userRepo"
	"inventory/internals/service/cryptoService"
	"inventory/internals/service/tokenService"
	"inventory/internals/service/validationService"
	"time"
)

type userSrv struct {
	repo       userRepo.UserRepository
	cryptoSrv  cryptoService.CryptoService
	token      tokenService.TokenService
	validation validationService.ValidationService
}

type UserService interface {
	CreateUser(req *userEntity.CreateUserReq) error
	GetUsers() ([]*userEntity.CreateUserRes, error)
	Login(req *userEntity.Login) (*userEntity.CreateUserRes, *string, error)
	GetUser(id string) (*userEntity.CreateUserRes, error)
	UpdateUser(id string, req *userEntity.UpdateUserReq) (*userEntity.CreateUserRes, error)
	DeleteUser(id string) error
}

func NewUserSrv(repo userRepo.UserRepository, validation validationService.ValidationService, cryptoSrv cryptoService.CryptoService, token tokenService.TokenService) UserService {
	return &userSrv{repo: repo, validation: validation, cryptoSrv: cryptoSrv, token: token}
}

// Create User
// @Summary	Create User
// @Description	Create A user
// @Tags	Auth
// @Accept	json
// @Produce	json
// @Param request body userEntity.CreateUserReq true "user details"
// @Success	200  {object} string
// @Router	/user [post]
func (t *userSrv) CreateUser(req *userEntity.CreateUserReq) error {

	if err := t.validation.Validate(req); err != nil {
		return err
	}

	_, err := t.repo.GetUserByEmail(req.Email)

	if err == nil {
		return fmt.Errorf("user exist already")
	}

	req.Password, _ = t.cryptoSrv.HashPassword(req.Password)
	req.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	req.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	return t.repo.CreateUser(req)
}

// Login
// @Summary	Login
// @Description	Login
// @Tags	Auth
// @Accept	json
// @Produce	json
// @Param	request	body userEntity.Login true "login details"
// @Success	200  {object} userEntity.CreateUserRes
// @Router	/user/login [post]
func (t *userSrv) Login(req *userEntity.Login) (*userEntity.CreateUserRes, *string, error) {
	user, err := t.repo.GetUserByEmail(req.Email)

	var res userEntity.CreateUserRes

	if err != nil {
		return nil, nil, fmt.Errorf("incorrect email or password")
	}

	hash := user.Password

	isMatch := t.cryptoSrv.ComparePassword(req.Password, hash)

	if !isMatch {
		return nil, nil, fmt.Errorf("incorrect password")
	}

	userJson, err := json.Marshal(user)

	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(userJson, &res)

	if err != nil {
		return nil, nil, err
	}

	token, refreshToken, err := t.token.CreateToken(user.Id.Hex(), user.Email)

	if err != nil {
		return nil, nil, err
	}

	var UpdateUser userEntity.UpdateUserReq

	UpdateUser.RefreshToken = &refreshToken

	_, err = t.repo.UpdateUser(user.Id.Hex(), &UpdateUser)

	if err != nil {
		return nil, nil, err
	}

	return &res, &token, nil
}

// Get Users
// @Summary	Get Users
// @Description	Get Users
// @Tags	User
// @Accept	json
// @Produce	json
// @Success	200  {object} []userEntity.CreateUserRes
// @Security ApiKeyAuth
// @Router	/user/ [get]
func (t *userSrv) GetUsers() ([]*userEntity.CreateUserRes, error) {
	users, err := t.repo.GetUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}

// Get User
// @Summary	Get User
// @Description	Get User
// @Tags	User
// @Accept	json
// @Produce	json
// @Param	request	path string true "user id"
// @Success	200  {object} userEntity.CreateUserRes
// @Security ApiKeyAuth
// @Router	/user/{id} [get]
func (t *userSrv) GetUser(id string) (*userEntity.CreateUserRes, error) {
	user, err := t.repo.GetUser(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// Update User
// @Summary	Update Users
// @Description	Update User
// @Tags	User
// @Accept	json
// @Produce	json
// @Param	request	path string true "user id"
// @Param	request	body userEntity.UpdateUserReq true "update details"
// @Success	200  {object} userEntity.CreateUserRes
// @Security ApiKeyAuth
// @Router	/user/{id} [patch]
func (t *userSrv) UpdateUser(id string, req *userEntity.UpdateUserReq) (*userEntity.CreateUserRes, error) {
	return t.repo.UpdateUser(id, req)
}

// Delete User
// @Summary	Delete Users
// @Description	Delete Users
// @Tags	User
// @Accept	json
// @Produce	json
// @Param	request	path string true "user id"
// @Success	200  {object} string
// @Security ApiKeyAuth
// @Router	/user/{id} [delete]
func (t *userSrv) DeleteUser(id string) error {
	return t.repo.DeleteUser(id)
}
