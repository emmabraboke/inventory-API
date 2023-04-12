package transactionRepo

import "inventory/internals/entity/transactionEntity"

type TransactionRepository interface {
	CreateTransaction(req *transactionEntity.Transaction) error
	GetTransactions() ([]*transactionEntity.Transaction, error)
	GetTransaction(id string) (*transactionEntity.Transaction, error)
	UpdateTransaction(id string, req *transactionEntity.UpdateTransaction) (*transactionEntity.Transaction, error)
	DeleteTransaction(id string) error
}
