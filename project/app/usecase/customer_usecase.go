package usecase

import (
	"context"

	"github.com/example/meteran-air/entity"
)

type CustomerUsecase interface {
	GetCustomerByID(ctx context.Context, id int64) (*entity.Customer, error)
	GetAllCustomers(ctx context.Context) ([]*entity.Customer, error)
	CreateCustomer(ctx context.Context, customer *entity.Customer) error
	UpdateCustomer(ctx context.Context, customer *entity.Customer) error
	DeleteCustomer(ctx context.Context, id int64) error
	SearchCustomers(ctx context.Context, query string) ([]*entity.Customer, error)
}

type customerUsecase struct {
	customerRepo entity.CustomerRepository
}

func NewCustomerUsecase(customerRepo entity.CustomerRepository) CustomerUsecase {
	return &customerUsecase{
		customerRepo: customerRepo,
	}
}

func (cu *customerUsecase) GetCustomerByID(ctx context.Context, id int64) (*entity.Customer, error) {
	return cu.customerRepo.GetByID(ctx, id)
}

func (cu *customerUsecase) GetAllCustomers(ctx context.Context) ([]*entity.Customer, error) {
	return cu.customerRepo.GetAll(ctx)
}

func (cu *customerUsecase) CreateCustomer(ctx context.Context, customer *entity.Customer) error {
	return cu.customerRepo.Create(ctx, customer)
}

func (cu *customerUsecase) UpdateCustomer(ctx context.Context, customer *entity.Customer) error {
	return cu.customerRepo.Update(ctx, customer)
}

func (cu *customerUsecase) DeleteCustomer(ctx context.Context, id int64) error {
	return cu.customerRepo.Delete(ctx, id)
}

func (cu *customerUsecase) SearchCustomers(ctx context.Context, query string) ([]*entity.Customer, error) {
	return cu.customerRepo.Search(ctx, query)
}
