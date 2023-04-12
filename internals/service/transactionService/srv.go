package transactionService

import (
	"inventory/internals/entity/invoiceEntity"
	"inventory/internals/entity/productEntity"
	"inventory/internals/entity/transactionEntity"
	"inventory/internals/repository/transactionRepo"
	"inventory/internals/service/invoiceService"
	"inventory/internals/service/paymentService"
	"inventory/internals/service/productService"
	"inventory/internals/service/validationService"
	"time"
)

type transactionSrv struct {
	repo       transactionRepo.TransactionRepository
	invoiceSrv invoiceService.InvoiceService
	validation validationService.ValidationService
	paymentSrv paymentService.PaymentService
	productSrv productService.ProductService
}

type TransactionService interface {
	CreateTransaction(req *transactionEntity.Transaction) (*transactionEntity.PayStackRes, error)
	GetTransactions() ([]*transactionEntity.Transaction, error)
	GetTransaction(id string) (*transactionEntity.Transaction, error)
	UpdateTransaction(id string, req *transactionEntity.UpdateTransaction) (*transactionEntity.Transaction, error)
	DeleteTransaction(id string) error
}

func NewtransactionSrv(repo transactionRepo.TransactionRepository, validation validationService.ValidationService, invoiceSrv invoiceService.InvoiceService, paymenySrv paymentService.PaymentService, productSrv productService.ProductService) TransactionService {
	return &transactionSrv{repo: repo, validation: validation, invoiceSrv: invoiceSrv, paymentSrv: paymenySrv, productSrv: productSrv}
}

func (t *transactionSrv) CreateTransaction(req *transactionEntity.Transaction) (*transactionEntity.PayStackRes, error) {

	if err := t.validation.Validate(req); err != nil {
		return nil, err
	}

	var paymentReq transactionEntity.PayStackReq

	//amount is divided by 100
	paymentReq.Amount = req.Amount * 100
	paymentReq.Email = req.Email

	paymentDetails, err := t.paymentSrv.CreatePayment(&paymentReq)

	if err != nil {
		return nil, err
	}

	req.Reference = paymentDetails.Data.Reference

	req.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	req.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	err = t.repo.CreateTransaction(req)

	if err != nil {
		return nil, err
	}

	return paymentDetails, nil
}

func (t *transactionSrv) GetTransactions() ([]*transactionEntity.Transaction, error) {
	transactions, err := t.repo.GetTransactions()

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *transactionSrv) GetTransaction(id string) (*transactionEntity.Transaction, error) {
	transaction, err := t.repo.GetTransaction(id)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *transactionSrv) UpdateTransaction(id string, req *transactionEntity.UpdateTransaction) (*transactionEntity.Transaction, error) {
	transaction, err := t.repo.UpdateTransaction(id, req)
	if err != nil {
		return nil, err
	}

	invoiceId := transaction.InvoiceId
	var update invoiceEntity.UpdateInvoice

	isPaid := true
	update.IsPaid = &isPaid

	invoice, err := t.invoiceSrv.UpdateInvoice(invoiceId.Hex(), &update)

	if err != nil {
		return nil, err
	}

	items := invoice.Item

	// update products
	for k := range items {
		productId := items[k].ProductId
		quantity := items[k].Quantity

		var updateProduct productEntity.UpdateProduct

		updateProduct.Quantity = &quantity

		_, err := t.productSrv.UpdateProductQuantity(productId.Hex(), &updateProduct)

		if err != nil {
			return nil, err
		}
	}

	return transaction, nil

}

func (t *transactionSrv) DeleteTransaction(id string) error {
	return t.repo.DeleteTransaction(id)
}
