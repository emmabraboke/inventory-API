package customerHandler

import (
	"inventory/internals/entity/customerEntity"
	"inventory/internals/service/customerService"
	"net/http"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	srv customerService.CustomerService
}

func NewcustomerHandler(srv customerService.CustomerService) *customerHandler {
	return &customerHandler{srv: srv}
}

func (t *customerHandler) CreateCustomer(c *gin.Context) {
	var req customerEntity.Customer
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = t.srv.CreateCustomer(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "account created successfully")
}

func (t *customerHandler) GetCustomers(c *gin.Context) {

	customer, err := t.srv.GetCustomers()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	

	c.JSON(http.StatusOK, customer)
}

func (t *customerHandler) GetCustomer(c *gin.Context) {
	id := c.Param("id")
	customer, err := t.srv.GetCustomer(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, customer)

}

func (t *customerHandler) UpdateCustomer(c *gin.Context) {
	var req customerEntity.UpdateCustomer

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Param("id")
	customer, err := t.srv.UpdateCustomer(id, &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (t *customerHandler) DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	err := t.srv.DeleteCustomer(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "customer deleted successfully")
}
