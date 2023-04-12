package routes

import (
	"inventory/cmd/handlers/saleHandler"
	"inventory/cmd/middlewares"
	"inventory/internals/service/saleService"
	"inventory/internals/service/tokenService"

	"github.com/gin-gonic/gin"
)

func SaleRoute(group *gin.RouterGroup, srv saleService.SaleService, tokenSrv tokenService.TokenService) {

	auth := middlewares.NewAuthMiddleware(tokenSrv)
	handler := saleHandler.NewSaleHandler(srv)
	route := group.Group("/sale")

	//middleware
	route.Use(auth.AuthMiddleware)
	route.POST("/", handler.CreateSale)
	route.GET("/", handler.GetSales)
	route.GET("/:id", handler.GetSale)
	route.PATCH("/id", handler.DeleteSale)
}
