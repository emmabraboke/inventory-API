package routes

import (
	"inventory/cmd/handlers/transactionHandler"
	"inventory/cmd/middlewares"
	"inventory/internals/service/tokenService"
	"inventory/internals/service/transactionService"

	"github.com/gin-gonic/gin"
)

func TransactionRoute(group *gin.RouterGroup, srv transactionService.TransactionService, tokenSrv tokenService.TokenService) {

	auth := middlewares.NewAuthMiddleware(tokenSrv)
	handler := transactionHandler.NewTransactionHandler(srv)
	route := group.Group("/transaction")

	//middleware
	route.Use(auth.AuthMiddleware)
	route.POST("/", handler.CreateTransaction)
	route.GET("/", handler.GetTransactions)
	route.GET("/:id", handler.GetTransaction)
	route.PATCH("/:id", handler.UpdateTransaction)
	route.DELETE("/id", handler.DeleteTransaction)
}
