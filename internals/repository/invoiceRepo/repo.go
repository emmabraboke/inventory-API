package invoiceRepo

import (
	"inventory/internals/entity/invoiceEntity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InvoiceRepository interface {
	CreateInvoice(req *invoiceEntity.Invoice) (*primitive.ObjectID, error)
	GetInvoices() ([]*invoiceEntity.Invoice, error)
	GetInvoice(id string) (*invoiceEntity.Invoice, error)
	UpdateInvoice(id string, req *invoiceEntity.UpdateInvoice) (*invoiceEntity.Invoice, error)
	DeleteInvoice(id string) error
}
