package repository

import (
	"fmt"

	"github.com/Ihpaz/golang-restapi-userfamily/entity"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Save(customer *entity.Customer) (*entity.Customer, error)
	FindAll() (*[]entity.Customer, error)
	FindCustomerByCstId(customer *entity.Customer, uid uint64) (*entity.Customer, error)
	UpdateACustomer(customer *entity.Customer, uid uint64) (*entity.Customer, error)
	DeleteACustomer(customer *entity.Customer, uid uint64) (int64, error)
}

type repo struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &repo{db}
}

func (r *repo) Save(customer *entity.Customer) (*entity.Customer, error) {
	var err error
	err = r.db.Create(&customer).Error
	if err != nil {
		return &entity.Customer{}, err
	}
	return customer, nil
}

func (r *repo) FindAll() (*[]entity.Customer, error) {
	var err error

	customers := []entity.Customer{}
	err = r.db.Find(&customers).Limit(100).Error

	if err != nil {
		return &[]entity.Customer{}, err
	}
	return &customers, nil
}

func (r *repo) FindCustomerByCstId(customer *entity.Customer, uid uint64) (*entity.Customer, error) {
	var err error
	err = r.db.Find(&customer).Where("cst_id = ?", uid).Error
	if err != nil {
		return &entity.Customer{}, err
	}

	return customer, err
}

func (r *repo) UpdateACustomer(customer *entity.Customer, uid uint64) (*entity.Customer, error) {
	var err error
	var query = r.db.Model(&entity.Customer{}).Where("cst_id = ?", uid).UpdateColumns(
		map[string]interface{}{
			"Nationality_id": customer.Nationality_id,
			"Cst_name":       customer.Cst_name,
			"Cst_email":      customer.Cst_email,
			"Cst_dob_date":   customer.Cst_dob_date,
			"Cst_phoneNum":   customer.Cst_phoneNum,
		},
	)

	if query.Error != nil {
		return &entity.Customer{}, query.Error
	}

	err = r.db.Find(&entity.Customer{}).Where("cst_id = ?", uid).Error
	if err != nil {
		return &entity.Customer{}, err
	}
	return customer, nil
}

func (r *repo) DeleteACustomer(customer *entity.Customer, uid uint64) (int64, error) {

	fmt.Println("masuk ko")
	var query = r.db.Where("cst_id = ?", uid).Delete(entity.Customer{})

	if query.Error != nil {
		return 0, query.Error
	}
	return query.RowsAffected, nil
}
