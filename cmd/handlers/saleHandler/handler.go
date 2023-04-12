package saleHandler

import (
	"inventory/internals/entity/saleEntity"
	"inventory/internals/service/saleService"
	"net/http"

	"github.com/gin-gonic/gin"
)

type saleHandler struct {
	srv saleService.SaleService
}

func NewSaleHandler(srv saleService.SaleService) *saleHandler {
	return &saleHandler{srv: srv}
}

func (t *saleHandler) CreateSale(c *gin.Context) {
	var req saleEntity.Sale
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = t.srv.CreateSale(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "account created successfully")
}

func (t *saleHandler) GetSales(c *gin.Context) {

	sale, err := t.srv.GetSales()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, sale)
}

func (t *saleHandler) GetSale(c *gin.Context) {
	id := c.Param("id")
	sale, err := t.srv.GetSale(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, sale)

}



func (t *saleHandler) DeleteSale(c *gin.Context) {
	id := c.Param("id")
	err := t.srv.DeleteSale(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "sale deleted successfully")
}
