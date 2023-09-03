package repository

import (
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
	err = r.db.Preload("Nationality").Find(&customers).Limit(100).Error
	if err != nil {
		return &[]entity.Customer{}, err
	}
	return &customers, nil
}

func (r *repo) FindCustomerByCstId(customer *entity.Customer, uid uint64) (*entity.Customer, error) {
	var err error
	err = r.db.Preload("FamilyList").Find(&customer, uid).Error
	if err != nil {
		return &entity.Customer{}, err
	}

	return customer, err
}

func (r *repo) UpdateACustomer(customer *entity.Customer, uid uint64) (*entity.Customer, error) {
	var err error
	var query = r.db.Model(&entity.Customer{ID: uint(uid)}).UpdateColumns(
		map[string]interface{}{
			"Nationality_id": customer.Nationality_id,
			"Cst_name":       customer.Cst_name,
			"Cst_email":      customer.Cst_email,
			"Cst_dob_date":   customer.Cst_dob_date,
			"Cst_phoneNum":   customer.Cst_phoneNum,
		},
	)
	query = r.db.Where("cst_id = ?", uid).Delete(entity.FamilyList{})

	for i := range customer.FamilyList {
		customer.FamilyList[i].Cst_id = int64(uid)
	}

	query = r.db.Create(customer.FamilyList)

	if query.Error != nil {
		return &entity.Customer{}, query.Error
	}
	err = r.db.Find(&entity.Customer{}, uid).Error
	if err != nil {
		return &entity.Customer{}, err
	}
	return customer, nil
}

func (r *repo) DeleteACustomer(customer *entity.Customer, uid uint64) (int64, error) {

	var query = r.db.Where("cst_id = ?", uid).Delete(entity.FamilyList{})
	query = r.db.Delete(&customer, uid)

	if query.Error != nil {
		return 0, query.Error
	}
	return query.RowsAffected, nil
}
