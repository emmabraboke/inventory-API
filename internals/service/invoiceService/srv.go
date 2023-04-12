package invoiceService

import (
	"inventory/internals/entity/invoiceEntity"
	"inventory/internals/entity/saleEntity"

	// "inventory/internals/entity/saleEntity"
	"inventory/internals/repository/invoiceRepo"
	"inventory/internals/service/saleService"
	"inventory/internals/service/validationService"
	"log"
	"time"
)

type invoiceSrv struct {
	repo       invoiceRepo.InvoiceRepository
	validation validationService.ValidationService
	saleSrv    saleService.SaleService
}

type InvoiceService interface {
	CreateInvoice(req *invoiceEntity.Invoice) error
	GetInvoices() ([]*invoiceEntity.Invoice, error)
	GetInvoice(id string) (*invoiceEntity.Invoice, error)
	UpdateInvoice(id string, req *invoiceEntity.UpdateInvoice) (*invoiceEntity.Invoice, error)
	DeleteInvoice(id string) error
}

func NewinvoiceSrv(repo invoiceRepo.InvoiceRepository, validation validationService.ValidationService, saleSrv saleService.SaleService) InvoiceService {
	return &invoiceSrv{repo: repo, validation: validation, saleSrv: saleSrv}
}

func (t *invoiceSrv) CreateInvoice(req *invoiceEntity.Invoice) error {

	if err := t.validation.Validate(req); err != nil {
		return err
	}

	req.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	req.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	invoice, err := t.repo.CreateInvoice(req)

	log.Println(req, invoice)

	if err != nil {
		return err
	}

	items := req.Item

	salesItem := make([]saleEntity.Sale, len(items))

	// id := invoice
	for k := range items {
		salesItem[k].InvoiceId = *invoice
		salesItem[k].CustomerId = req.CustomerId
		salesItem[k].ProductId = items[k].ProductId
		salesItem[k].Name = items[k].Name
		salesItem[k].Price = items[k].Price
		salesItem[k].Quantity = items[k].Quantity
		err = t.saleSrv.CreateSale(&salesItem[k])
		if err != nil {
			return nil
		}
	}

	return nil
}

func (t *invoiceSrv) GetInvoices() ([]*invoiceEntity.Invoice, error) {
	invoices, err := t.repo.GetInvoices()

	if err != nil {
		return nil, err
	}

	return invoices, nil
}

func (t *invoiceSrv) GetInvoice(id string) (*invoiceEntity.Invoice, error) {
	invoice, err := t.repo.GetInvoice(id)

	if err != nil {
		return nil, err
	}

	return invoice, nil
}

func (t *invoiceSrv) UpdateInvoice(id string, req *invoiceEntity.UpdateInvoice) (*invoiceEntity.Invoice, error) {
	return t.repo.UpdateInvoice(id, req)
}

func (t *invoiceSrv) DeleteInvoice(id string) error {
	return t.repo.DeleteInvoice(id)
}
