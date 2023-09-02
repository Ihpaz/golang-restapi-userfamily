package config

import (
	"log"

	"github.com/Ihpaz/golang-restapi-userfamily/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// dbURL := "postgres://pg:pass@localhost:5432/bookingapp"

	dsn := "host=localhost user=postgres password=admin dbname=bookingapp port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&entity.Customer{})
	db.AutoMigrate(&entity.Nationality{})
	db.AutoMigrate(&entity.Family_List{})

	return db
}
