package productService

import (
	"inventory/internals/entity/productEntity"
	"inventory/internals/repository/productRepo"
	"inventory/internals/service/validationService"
	"time"
)

type productSrv struct {
	repo       productRepo.ProductRepository
	validation validationService.ValidationService
}

type ProductService interface {
	CreateProduct(req *productEntity.Product) error
	GetProducts() ([]*productEntity.Product, error)
	GetProduct(id string) (*productEntity.Product, error)
	UpdateProduct(id string, req *productEntity.UpdateProduct) (*productEntity.Product, error)
	UpdateProductQuantity(id string, req *productEntity.UpdateProduct) (*productEntity.Product, error)
	DeleteProduct(id string) error
}

func NewproductSrv(repo productRepo.ProductRepository, validation validationService.ValidationService) ProductService {
	return &productSrv{repo: repo, validation: validation}
}

func (t *productSrv) CreateProduct(req *productEntity.Product) error {

	if err := t.validation.Validate(req); err != nil {
		return err
	}

	req.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	req.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	return t.repo.CreateProduct(req)
}

func (t *productSrv) GetProducts() ([]*productEntity.Product, error) {
	products, err := t.repo.GetProducts()

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (t *productSrv) GetProduct(id string) (*productEntity.Product, error) {
	product, err := t.repo.GetProduct(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (t *productSrv) UpdateProduct(id string, req *productEntity.UpdateProduct) (*productEntity.Product, error) {
	return t.repo.UpdateProduct(id, req)
}

func (t *productSrv) UpdateProductQuantity(id string, req *productEntity.UpdateProduct) (*productEntity.Product, error) {
	return t.repo.UpdateProductQuantity(id, req)
}

func (t *productSrv) DeleteProduct(id string) error {
	return t.repo.DeleteProduct(id)
}
