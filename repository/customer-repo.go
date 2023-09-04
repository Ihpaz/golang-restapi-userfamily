package repository

import (
	"github.com/Ihpaz/golang-restapi-userfamily/entity"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Save(customer *entity.Customer) (*entity.Customer, error)
	FindAll() ([]entity.Customer, error)
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

func (r *repo) FindAll() ([]entity.Customer, error) {
	var err error
	customers := []entity.Customer{}
	err = r.db.Preload("Nationality").Find(&customers).Limit(100).Error
	if err != nil {
		return []entity.Customer{}, err
	}
	return customers, nil
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
	tx := r.db.Begin()

	var query = tx.Model(&entity.Customer{ID: uint(uid)}).UpdateColumns(
		map[string]interface{}{
			"Nationality_id": customer.Nationality_id,
			"Cst_name":       customer.Cst_name,
			"Cst_email":      customer.Cst_email,
			"Cst_dob_date":   customer.Cst_dob_date,
			"Cst_phoneNum":   customer.Cst_phoneNum,
		},
	)

	var query2 = tx.Where("cst_id = ?", uid).Delete(entity.FamilyList{})
	if query2.Error != nil {
		tx.Rollback()
		return &entity.Customer{}, query.Error
	}

	for i := range customer.FamilyList {
		customer.FamilyList[i].Cst_id = int64(uid)
	}

	var query3 = tx.Create(customer.FamilyList)
	if query3.Error != nil {
		tx.Rollback()
		return &entity.Customer{}, query.Error
	}

	tx.Commit()

	err = r.db.Find(&entity.Customer{}, uid).Error
	if err != nil {
		return &entity.Customer{}, err
	}
	return customer, nil
}

func (r *repo) DeleteACustomer(customer *entity.Customer, uid uint64) (int64, error) {

	tx := r.db.Begin()

	var query1 = tx.Where("cst_id = ?", uid).Delete(entity.FamilyList{})
	if query1.Error != nil || query1.RowsAffected == 0 {
		return 0, query1.Error
	}

	var query2 = tx.Delete(&customer, uid)
	if query2.Error != nil || query2.RowsAffected == 0 {
		tx.Rollback()
		return 0, query2.Error
	}

	tx.Commit()
	return query2.RowsAffected, nil
}
