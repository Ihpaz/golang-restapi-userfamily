package entity

import (
	"time"
)

type Customer struct {
	cst_id         int64 `gorm:"primaryKey"`
	nationality_id int64
	cst_name       string
	cst_dob_date   time.Time
	cst_phoneNum   string
	cst_email      string
}
