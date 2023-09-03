package config

import (
	"log"

	"github.com/Ihpaz/golang-restapi-userfamily/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {

	nationalities := []entity.Nationality{
		{Nationality_name: "Indonesia", Nationality_code: "ID"},
		{Nationality_name: "Malaysia", Nationality_code: "MY"},
		{Nationality_name: "Singapore", Nationality_code: "SG"},
	}

	var result = db.Find(&entity.Nationality{})

	if result.RowsAffected == 0 {
		for _, nationality := range nationalities {
			db.Create(&nationality)
		}
	}

}

func Init() *gorm.DB {
	// dbURL := "postgres://pg:pass@localhost:5432/bookingapp"

	dsn := "host=localhost user=postgres password=admin dbname=bookingapp port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&entity.Nationality{}, &entity.Customer{}, &entity.FamilyList{})

	SeedUsers(db)

	return db
}
