package customerService

import (
	"fmt"
	"inventory/internals/entity/customerEntity"
	"inventory/internals/repository/customerRepo"
	"inventory/internals/service/validationService"
	"time"
)

type customerSrv struct {
	repo       customerRepo.CustomerRepository
	validation validationService.ValidationService
}

type CustomerService interface {
	CreateCustomer(req *customerEntity.Customer) error
	GetCustomers() ([]*customerEntity.Customer, error)
	GetCustomer(id string) (*customerEntity.Customer, error)
	UpdateCustomer(id string, req *customerEntity.UpdateCustomer) (*customerEntity.Customer, error)
	DeleteCustomer(id string) error
}

func NewCustomerSrv(repo customerRepo.CustomerRepository, validation validationService.ValidationService) CustomerService {
	return &customerSrv{repo: repo, validation: validation}
}

func (t *customerSrv) CreateCustomer(req *customerEntity.Customer) error {

	if err := t.validation.Validate(req); err != nil {
		return err
	}
	
	
	_, err := t.repo.GetCustomerByEmail(req.Email)

	if err == nil {
		return fmt.Errorf("customer exist already")
	}



	req.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	req.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	return t.repo.CreateCustomer(req)

}

func (t *customerSrv) GetCustomers() ([]*customerEntity.Customer, error) {
	Customers, err := t.repo.GetCustomers()

	if err != nil {
		return nil, err
	}

	return Customers, nil
}

func (t *customerSrv) GetCustomer(id string) (*customerEntity.Customer, error) {
	Customer, err := t.repo.GetCustomer(id)

	if err != nil {
		return nil, err
	}

	return Customer, nil
}

func (t *customerSrv) UpdateCustomer(id string, req *customerEntity.UpdateCustomer) (*customerEntity.Customer, error) {
	return t.repo.UpdateCustomer(id, req)
}

func (t *customerSrv) DeleteCustomer(id string) error {
	return t.repo.DeleteCustomer(id)
}
