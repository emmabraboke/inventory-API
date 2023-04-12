package routes

import (
	"inventory/cmd/handlers/invoiceHandler"
	"inventory/cmd/middlewares"
	"inventory/internals/service/invoiceService"
	"inventory/internals/service/tokenService"

	"github.com/gin-gonic/gin"
)

func InvoiceRoute(group *gin.RouterGroup, srv invoiceService.InvoiceService, tokenSrv tokenService.TokenService) {

	auth := middlewares.NewAuthMiddleware(tokenSrv)
	handler := invoiceHandler.NewInvoiceHandler(srv)
	route := group.Group("/invoice")

	route.Use(auth.AuthMiddleware)
	route.POST("/", handler.CreateInvoice)
	route.GET("/", handler.GetInvoices)
	route.GET("/:id", handler.GetInvoice)
	route.PATCH("/:id", handler.UpdateInvoice)
	route.PATCH("/id", handler.DeleteInvoice)
}
