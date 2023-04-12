package routes

import (
	"inventory/cmd/handlers/userHandler"
	 "inventory/cmd/middlewares"
	"inventory/internals/service/tokenService"
	"inventory/internals/service/userService"

	"github.com/gin-gonic/gin"
)

func UserRoute(group *gin.RouterGroup, srv userService.UserService, tokenSrv tokenService.TokenService) {

	auth := middlewares.NewAuthMiddleware(tokenSrv)

	handler := userHandler.NewUserHandler(srv)

	route := group.Group("/user")
	route.POST("/", handler.SignUp)
	route.POST("/login", handler.Login)

	// auth middleware
	route.Use(auth.AuthMiddleware)
	route.GET("/", handler.GetUsers)
	route.GET("/:id", handler.GetUser)
	route.PATCH("/:id", handler.UpdateUser)
	route.DELETE("/id", handler.DeleteUser)

}
