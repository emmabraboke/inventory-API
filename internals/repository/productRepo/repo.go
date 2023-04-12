package productRepo

import "inventory/internals/entity/productEntity"

type ProductRepository interface {
	CreateProduct(req *productEntity.Product) error
	GetProducts() ([]*productEntity.Product, error)
	GetProduct(id string) (*productEntity.Product, error)
	UpdateProduct(id string, req *productEntity.UpdateProduct) (*productEntity.Product, error)
	UpdateProductQuantity(id string, req *productEntity.UpdateProduct) (*productEntity.Product, error)
	DeleteProduct(id string) error
}
