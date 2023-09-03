package repository

import (
	"github.com/Ihpaz/golang-restapi-userfamily/entity"
	"github.com/stretchr/testify/mock"
)

type MockCustomerRepository interface {
	Save(customer *entity.Customer) (*entity.Customer, error)
	FindAll() ([]entity.Customer, error)
	// FindCustomerByCstId(customer *entity.Customer, uid uint64) (*entity.Customer, error)
	// UpdateACustomer(customer *entity.Customer, uid uint64) (*entity.Customer, error)
	// DeleteACustomer(customer *entity.Customer, uid uint64) (int64, error)
}

type repoCustomerRepository struct {
	mock.Mock
}

func NewMockCustomerRepository(mock mock.Mock) MockCustomerRepository {
	return &repoCustomerRepository{mock}
}

func (mock *repoCustomerRepository) Save(post *entity.Customer) (*entity.Customer, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Customer), args.Error(1)
}

func (mock *repoCustomerRepository) FindAll() ([]entity.Customer, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Customer), args.Error(1)
}
