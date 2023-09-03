package entity

import "time"

type FamilyList struct {
	Fl_id       int64 `gorm:"primaryKey"`
	Cst_id      int64
	Fl_relation string
	Fl_name     string
	Fl_dob_date time.Time
}
