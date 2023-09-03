package entity

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type CustomTime struct {
	time.Time
}

func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return
}

func (ct CustomTime) Value() (driver.Value, error) {
	return ct.Time, nil
}

func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		ct.Time = time.Time{}
		return nil
	}
	parsedTime, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("Invalid type for CustomTime")
	}
	ct.Time = parsedTime
	return nil
}

type Customer struct {
	Cst_id         int64 `gorm:"primaryKey"`
	Nationality_id int64
	Cst_name       string
	Cst_dob_date   CustomTime `json:"Cst_dob_date"`
	Cst_phoneNum   string
	Cst_email      string
	Nationality    Nationality
	FamilyList     []FamilyList `gorm:"foreignKey:Cst_id"`
}
