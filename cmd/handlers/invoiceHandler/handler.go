package invoiceHandler

import (
	"inventory/internals/entity/invoiceEntity"
	"inventory/internals/service/invoiceService"
	"net/http"

	"github.com/gin-gonic/gin"
)

type invoiceHandler struct {
	srv invoiceService.InvoiceService
}

func NewInvoiceHandler(srv invoiceService.InvoiceService) *invoiceHandler {
	return &invoiceHandler{srv: srv}
}

func (t *invoiceHandler) CreateInvoice(c *gin.Context) {
	var req invoiceEntity.Invoice
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = t.srv.CreateInvoice(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "invoice created successfully")
}

func (t *invoiceHandler) GetInvoices(c *gin.Context) {

	invoice, err := t.srv.GetInvoices()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, invoice)
}

func (t *invoiceHandler) GetInvoice(c *gin.Context) {
	id := c.Param("id")
	invoice, err := t.srv.GetInvoice(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, invoice)

}

func (t *invoiceHandler) UpdateInvoice(c *gin.Context) {
	var req invoiceEntity.UpdateInvoice

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Param("id")
	invoice, err := t.srv.UpdateInvoice(id, &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, invoice)
}

func (t *invoiceHandler) DeleteInvoice(c *gin.Context) {
	id := c.Param("id")
	err := t.srv.DeleteInvoice(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "invoice deleted successfully")
}
