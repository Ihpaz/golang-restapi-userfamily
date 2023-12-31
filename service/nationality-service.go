package service

import (
	"errors"

	"github.com/Ihpaz/golang-restapi-userfamily/entity"
	"github.com/Ihpaz/golang-restapi-userfamily/repository"
)

type NationalityService interface {
	Validate(nationality *entity.Nationality) error
	Create(nationality *entity.Nationality) (*entity.Nationality, error)
	FindAll() (*[]entity.Nationality, error)
}

type servicenationality struct{}

var (
	reponationality repository.NationalityRepository
)

func NewNationalityService(repository repository.NationalityRepository) NationalityService {
	reponationality = repository
	return &servicenationality{}
}

func (*servicenationality) Validate(nationality *entity.Nationality) error {
	if nationality == nil {
		err := errors.New("The nationality is empty")
		return err
	}
	if nationality.Nationality_name == "" {
		err := errors.New("The nationality name is empty")
		return err
	}
	if nationality.Nationality_code == "" {
		err := errors.New("The nationality empty")
		return err
	}
	return nil
}

func (*servicenationality) Create(nationality *entity.Nationality) (*entity.Nationality, error) {
	return reponationality.Save(nationality)
}

func (*servicenationality) FindAll() (*[]entity.Nationality, error) {
	return reponationality.FindAll()
}
