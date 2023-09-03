package service

import (
	"testing"

	"github.com/Ihpaz/golang-restapi-userfamily/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCustomerRepository struct {
	mock.Mock
}

func (mock *MockCustomerRepository) Save(customer *entity.Customer) (*entity.Customer, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Customer), args.Error(1)
}

func (mock *MockCustomerRepository) FindAll() ([]entity.Customer, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Customer), args.Error(1)
}

func (mock *MockCustomerRepository) FindCustomerByCstId(customer *entity.Customer, uid uint64) (*entity.Customer, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Customer), args.Error(1)
}

func (mock *MockCustomerRepository) UpdateACustomer(customer *entity.Customer, uid uint64) (*entity.Customer, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Customer), args.Error(1)
}

func (mock *MockCustomerRepository) DeleteACustomer(customer *entity.Customer, uid uint64) (int64, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(int64), args.Error(1)
}

func TestValidateEmptyCustomer(t *testing.T) {
	testService := NewCustomerService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "The customer is empty", err.Error())
}

func TestValidateEmptyCustomerName(t *testing.T) {

	defaultFamily1 := entity.FamilyList{
		Fl_name:     "defaultFamilyName1",
		Fl_dob_date: entity.CustomTimeFl{},
	}

	customer := entity.Customer{
		Cst_name:       "",
		Nationality_id: 1,
		Cst_dob_date:   entity.CustomTime{},
		FamilyList:     []entity.FamilyList{defaultFamily1}}

	testService := NewCustomerService(nil)
	err := testService.Validate(&customer)

	assert.NotNil(t, err)
	assert.Equal(t, "The customer name is empty", err.Error())
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockCustomerRepository)
	defaultFamily1 := entity.FamilyList{
		Fl_name:     "defaultFamilyName1",
		Fl_dob_date: entity.CustomTimeFl{},
	}
	customer := entity.Customer{
		Cst_name:       "Ihpaz",
		Nationality_id: 1,
		Cst_email:      "ihfazm@gmail.com",
		Cst_dob_date:   entity.CustomTime{},
		FamilyList:     []entity.FamilyList{defaultFamily1}}

	mockRepo.On("FindAll").Return([]entity.Customer{customer}, nil)

	testService := NewCustomerService(mockRepo)

	result, _ := testService.FindAll()

	mockRepo.AssertExpectations(t)

	assert.Equal(t, "Ihpaz", result[0].Cst_name)
	assert.Equal(t, "ihfazm@gmail.com", result[0].Cst_email)
	assert.Equal(t, int64(1), result[0].Nationality_id)
}

func TestSave(t *testing.T) {
	mockRepo := new(MockCustomerRepository)
	defaultFamily1 := entity.FamilyList{
		Fl_name:     "defaultFamilyName1",
		Fl_dob_date: entity.CustomTimeFl{},
	}
	customer := entity.Customer{
		Cst_name:       "Ihpaz",
		Nationality_id: 1,
		Cst_email:      "ihfazm@gmail.com",
		Cst_dob_date:   entity.CustomTime{},
		FamilyList:     []entity.FamilyList{defaultFamily1}}

	mockRepo.On("Save").Return(&customer, nil)

	testService := NewCustomerService(mockRepo)

	result, err := testService.Create(&customer)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, "Ihpaz", result.Cst_name)
	assert.Equal(t, "ihfazm@gmail.com", result.Cst_email)
	assert.Equal(t, int64(1), result.Nationality_id)
	assert.Nil(t, err)
}

func TestFindCustomerByCstId(t *testing.T) {
	mockRepo := new(MockCustomerRepository)
	defaultFamily1 := entity.FamilyList{
		Fl_name:     "defaultFamilyName1",
		Fl_dob_date: entity.CustomTimeFl{},
	}
	customer := entity.Customer{
		ID:             6,
		Cst_name:       "Ihpaz",
		Nationality_id: 1,
		Cst_email:      "ihfazm@gmail.com",
		Cst_dob_date:   entity.CustomTime{},
		FamilyList:     []entity.FamilyList{defaultFamily1}}

	mockRepo.On("FindCustomerByCstId").Return(&customer, nil)

	testService := NewCustomerService(mockRepo)

	result, err := testService.FindCustomerByCstId(&customer, 6)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, uint(6), result.ID)
	assert.Equal(t, "Ihpaz", result.Cst_name)
	assert.Equal(t, "ihfazm@gmail.com", result.Cst_email)
	assert.Nil(t, err)
}

func TestUpdateACustomer(t *testing.T) {
	mockRepo := new(MockCustomerRepository)
	defaultFamily1 := entity.FamilyList{
		Fl_name:     "defaultFamilyName1",
		Fl_dob_date: entity.CustomTimeFl{},
	}
	customer := entity.Customer{
		ID:             6,
		Cst_name:       "Ihpaz",
		Nationality_id: 1,
		Cst_email:      "ihfazm@gmail.com",
		Cst_dob_date:   entity.CustomTime{},
		FamilyList:     []entity.FamilyList{defaultFamily1}}

	mockRepo.On("UpdateACustomer").Return(&customer, nil)

	testService := NewCustomerService(mockRepo)

	result, err := testService.UpdateACustomer(&customer, 6)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, uint(6), result.ID)
	assert.Equal(t, "Ihpaz", result.Cst_name)
	assert.Equal(t, "ihfazm@gmail.com", result.Cst_email)
	assert.Nil(t, err)
}

func TestDeleteACustomer(t *testing.T) {
	mockRepo := new(MockCustomerRepository)

	customer := entity.Customer{}

	mockRepo.On("DeleteACustomer").Return(int64(1), nil)

	testService := NewCustomerService(mockRepo)

	result, err := testService.DeleteACustomer(&customer, 6)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, int64(1), result)
	assert.Nil(t, err)
}
