package entity

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type CustomTimeFl struct {
	time.Time
}

func (t *CustomTimeFl) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return
}

func (ct CustomTimeFl) Value() (driver.Value, error) {
	return ct.Time, nil
}

func (ct *CustomTimeFl) Scan(value interface{}) error {
	if value == nil {
		ct.Time = time.Time{}
		return nil
	}
	parsedTime, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("Invalid type for CustomTimeFl")
	}
	ct.Time = parsedTime
	return nil
}

type FamilyList struct {
	ID          uint
	Cst_id      int64
	Fl_relation string
	Fl_name     string
	Fl_dob_date CustomTimeFl
}
