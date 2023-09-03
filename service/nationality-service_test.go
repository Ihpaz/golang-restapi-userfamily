package service

import (
	"testing"

	"github.com/Ihpaz/golang-restapi-userfamily/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MocknationalityRepository struct {
	mock.Mock
}

func TestValidateEmptyNationality(t *testing.T) {
	testService := NewNationalityService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "The nationality is empty", err.Error())
}

func TestValidateEmptyNationalityName(t *testing.T) {

	nationality := entity.Nationality{
		Nationality_name: "",
		Nationality_code: "abc",
	}

	testService := NewNationalityService(nil)
	err := testService.Validate(&nationality)

	assert.NotNil(t, err)
	assert.Equal(t, "The nationality name is empty", err.Error())
}
