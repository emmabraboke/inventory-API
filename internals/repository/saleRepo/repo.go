package saleRepo

import "inventory/internals/entity/saleEntity"

type SaleRepository interface {
	CreateSale(req *saleEntity.Sale) error
	GetSales() ([]*saleEntity.Sale, error)
	GetSale(id string) (*saleEntity.Sale, error)
	DeleteSale(id string) error
}
