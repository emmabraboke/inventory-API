package routes

import (
	"inventory/cmd/handlers/customerHandler"
	"inventory/cmd/middlewares"
	"inventory/internals/service/customerService"
	"inventory/internals/service/tokenService"

	"github.com/gin-gonic/gin"
)

func CustomerRoute(group *gin.RouterGroup, srv customerService.CustomerService, tokenSrv tokenService.TokenService) {
	auth := middlewares.NewAuthMiddleware(tokenSrv)
	handler := customerHandler.NewcustomerHandler(srv)
	route := group.Group("/customer")

	//middleware
	route.Use(auth.AuthMiddleware)
	route.POST("/", handler.CreateCustomer)
	route.GET("/", handler.GetCustomers)
	route.GET("/:id", handler.GetCustomer)
	route.PATCH("/:id", handler.UpdateCustomer)
	route.PATCH("/id", handler.DeleteCustomer)

}
