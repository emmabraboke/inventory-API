package routes

import (
	"inventory/cmd/handlers/productHandler"
	"inventory/cmd/middlewares"
	"inventory/internals/service/productService"
	"inventory/internals/service/tokenService"

	"github.com/gin-gonic/gin"
)

func ProductRoute(group *gin.RouterGroup, srv productService.ProductService, tokenSrv tokenService.TokenService) {

	auth := middlewares.NewAuthMiddleware(tokenSrv)
	handler := productHandler.NewProductHandler(srv)
	route := group.Group("/product")

	//middleware
	route.Use((auth.AuthMiddleware))
	route.POST("/", handler.CreateProduct)
	route.GET("/", handler.GetProducts)
	route.GET("/:id", handler.GetProduct)
	route.PATCH("/:id", handler.UpdateProduct)
	route.PATCH("/id", handler.DeleteProduct)
}
