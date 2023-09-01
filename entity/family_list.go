package entity

import "time"

type Family_List struct {
	fl_id       int64 `gorm:"primaryKey"`
	cst_id      int64
	fl_relation string
	fl_name     string
	fl_dob_date time.Time
}
