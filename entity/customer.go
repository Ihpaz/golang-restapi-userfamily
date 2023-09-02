package entity

import "time"

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

type Customer struct {
	Cst_id         int64 `gorm:"primaryKey"`
	Nationality_id int64
	Cst_name       string
	Cst_dob_date   CustomTime `json:"Cst_dob_date"`
	Cst_phoneNum   string
	Cst_email      string
}
