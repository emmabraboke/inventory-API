package transactionHandler

import (
	"inventory/internals/entity/transactionEntity"
	"inventory/internals/service/transactionService"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	srv transactionService.TransactionService
}

func NewTransactionHandler(srv transactionService.TransactionService) *transactionHandler {
	return &transactionHandler{srv: srv}
}

func (t *transactionHandler) CreateTransaction(c *gin.Context) {
	var req transactionEntity.Transaction
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}


	payment, err := t.srv.CreateTransaction(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, payment)
}

func (t *transactionHandler) GetTransactions(c *gin.Context) {

	transaction, err := t.srv.GetTransactions()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, transaction)
}

func (t *transactionHandler) GetTransaction(c *gin.Context) {
	id := c.Param("id")
	transaction, err := t.srv.GetTransaction(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, transaction)

}

func (t *transactionHandler) UpdateTransaction(c *gin.Context) {
	var req transactionEntity.UpdateTransaction

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := c.Param("id")
	transaction, err := t.srv.UpdateTransaction(id, &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (t *transactionHandler) DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	err := t.srv.DeleteTransaction(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "transaction deleted successfully")
}
