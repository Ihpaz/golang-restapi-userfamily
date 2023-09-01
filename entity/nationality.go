package entity

type Nationality struct {
	nationality_id   int64 `gorm:"primaryKey"`
	nationality_name string
	nationality_code string
}
