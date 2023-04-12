package customerRepo

import customerEntity "inventory/internals/entity/customerEntity"

type CustomerRepository interface {
	CreateCustomer(req *customerEntity.Customer) error
	GetCustomers() ([]*customerEntity.Customer, error)
	GetCustomerByEmail(email string) (*customerEntity.Customer, error)
	GetCustomer(id string) (*customerEntity.Customer, error)
	UpdateCustomer(id string, req *customerEntity.UpdateCustomer) (*customerEntity.Customer, error)
	DeleteCustomer(id string) error
}
