package productHandler

import (
	"inventory/internals/entity/productEntity"
	"inventory/internals/service/productService"
	"net/http"
	"github.com/gin-gonic/gin"
)

type productHandler struct {
	srv productService.ProductService
}

func NewProductHandler(srv productService.ProductService) *productHandler {
	return &productHandler{srv: srv}
}

func (t *productHandler) CreateProduct(c *gin.Context) {
	var req productEntity.Product
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = t.srv.CreateProduct(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "account created successfully")
}

func (t *productHandler) GetProducts(c *gin.Context) {

	product, err := t.srv.GetProducts()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, product)
}

func (t *productHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := t.srv.GetProduct(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, product)

}

func (t *productHandler) UpdateProduct(c *gin.Context) {
	var req productEntity.UpdateProduct

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Param("id")
	product, err := t.srv.UpdateProduct(id, &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}

func (t *productHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	err := t.srv.DeleteProduct(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "product deleted successfully")
}
