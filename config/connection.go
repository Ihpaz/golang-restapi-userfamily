package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Ihpaz/golang-restapi-userfamily/entity"
	"github.com/joho/godotenv"
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

func LoadEnv() {
	rootDir, err := os.Getwd()
	envFilePath := filepath.Join(rootDir, ".env")
	err = godotenv.Load(envFilePath)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func Init() *gorm.DB {
	LoadEnv()

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&entity.Nationality{}, &entity.Customer{}, &entity.FamilyList{})

	if os.Getenv("APP_ENV") == "DEV" {
		SeedUsers(db)
	}

	return db
}
