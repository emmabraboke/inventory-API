package saleService

import (
	"inventory/internals/entity/saleEntity"
	"inventory/internals/repository/saleRepo"
	"inventory/internals/service/validationService"
	"time"
)

type saleSrv struct {
	repo       saleRepo.SaleRepository
	validation validationService.ValidationService
}

type SaleService interface {
	CreateSale(req *saleEntity.Sale) error
	GetSales() ([]*saleEntity.Sale, error)
	GetSale(id string) (*saleEntity.Sale, error)
	DeleteSale(id string) error
}

func NewsaleSrv(repo saleRepo.SaleRepository, validation validationService.ValidationService) SaleService {
	return &saleSrv{repo: repo, validation: validation}
}

func (t *saleSrv) CreateSale(req *saleEntity.Sale) error {

	if err := t.validation.Validate(req); err != nil {
		return err
	}

	req.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	req.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	return t.repo.CreateSale(req)
}

func (t *saleSrv) GetSales() ([]*saleEntity.Sale, error) {
	sales, err := t.repo.GetSales()

	if err != nil {
		return nil, err
	}

	return sales, nil
}

func (t *saleSrv) GetSale(id string) (*saleEntity.Sale, error) {
	sale, err := t.repo.GetSale(id)

	if err != nil {
		return nil, err
	}

	return sale, nil
}

func (t *saleSrv) DeleteSale(id string) error {
	return t.repo.DeleteSale(id)
}
