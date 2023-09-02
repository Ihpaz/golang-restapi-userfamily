package repository

import (
	"github.com/Ihpaz/golang-restapi-userfamily/entity"
	"gorm.io/gorm"
)

type NationalityRepository interface {
	Save(nationality *entity.Nationality) (*entity.Nationality, error)
	FindAll() (*[]entity.Nationality, error)
}

type reponationality struct {
	db *gorm.DB
}

func NewNationalityRepository(db *gorm.DB) NationalityRepository {
	return &reponationality{db}
}

func (r *reponationality) Save(nationality *entity.Nationality) (*entity.Nationality, error) {
	var err error
	err = r.db.Create(&nationality).Error
	if err != nil {
		return &entity.Nationality{}, err
	}
	return nationality, nil
}

func (r *reponationality) FindAll() (*[]entity.Nationality, error) {
	var err error

	naationalities := []entity.Nationality{}
	err = r.db.Find(&naationalities).Limit(100).Error

	if err != nil {
		return &[]entity.Nationality{}, err
	}
	return &naationalities, nil
}
