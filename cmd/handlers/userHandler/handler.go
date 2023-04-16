package userHandler

import (
	"inventory/internals/entity/responseEntity"
	"inventory/internals/entity/userEntity"
	"inventory/internals/service/userService"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	srv userService.UserService
}

func NewUserHandler(srv userService.UserService) *userHandler {
	return &userHandler{srv: srv}
}

func (t *userHandler) SignUp(c *gin.Context) {
	var req userEntity.CreateUserReq
	var file userEntity.ImageFile

	image, _, err := c.Request.FormFile("file")

	if err != nil {
		log.Println("no file provided")
	}

	err = c.ShouldBind(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	file.File = image

	err = t.srv.CreateUser(&req, file)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "account created successfully")
}

func (t *userHandler) Login(c *gin.Context) {
	var req userEntity.Login

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, token, err := t.srv.Login(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	log.Println(responseEntity.LoginResponse(http.StatusOK, "login successful", token, res))

	c.JSON(http.StatusOK, responseEntity.LoginResponse(http.StatusOK, "login successful", token, res))

}

func (t *userHandler) GetUsers(c *gin.Context) {

	user, err := t.srv.GetUsers()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, user)
}

func (t *userHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := t.srv.GetUser(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, user)

}

func (t *userHandler) UpdateUser(c *gin.Context) {
	var req userEntity.UpdateUserReq

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Param("id")
	user, err := t.srv.UpdateUser(id, &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (t *userHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	user := t.srv.DeleteUser(id)
	c.JSON(http.StatusOK, user)

}
