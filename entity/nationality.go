package entity

type Nationality struct {
	Nationality_id   int64 `gorm:"primaryKey"`
	Nationality_name string
	Nationality_code string
}
