package service

// import "testing"
import (
	"testing"

	"github.com/Ihpaz/golang-restapi-userfamily/entity"
	"github.com/stretchr/testify/assert"
)

var mockCustomerRepository = &repository.(Mock)

func TestValidateEmptyCustomer(t *testing.T) {
	testService := NewCustomerService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "he customer is empty", err.Error())
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
	mockRepo := new(mockCustomerRepository)

	var identifier int64 = 1

	defaultFamily1 := entity.FamilyList{
		Fl_name:     "defaultFamilyName1",
		Fl_dob_date: entity.CustomTimeFl{},
	}
	customer := entity.Customer{
		Cst_name:       "Ihpaz",
		Nationality_id: 1,
		Cst_phoneNum:   "ihfazm@gmail.com",
		Cst_dob_date:   entity.CustomTime{},
		FamilyList:     []entity.FamilyList{defaultFamily1}}
	// Setup expectations
	mockRepo.On("FindAll").Return([]entity.Customer{customer}, nil)

	testService := NewCustomerService(mockRepo)

	result, _ := testService.FindAll()

	//Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, identifier, resultCst_name)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)
}
