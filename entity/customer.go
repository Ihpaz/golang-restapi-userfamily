package entity

import (
	"time"
)

type Customer struct {
	Cst_id         int64 `gorm:"primaryKey"`
	Nationality_id int64
	Cst_name       string
	Cst_dob_date   time.Time
	Cst_phoneNum   string
	Cst_email      string
}
