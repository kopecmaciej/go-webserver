package lib

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dns = os.Getenv("DATABASE_URL")

var DB *gorm.DB

func Open() *gorm.DB {
	if len(dns) == 0 {
		dns = "postgresql://postgres:password@localhost:5432/web-app?sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection to db properly initiated")

	DB = db

	return db
}

func AutoMigrate(models ...interface{}) {
	err := DB.AutoMigrate(models...)
	if err != nil {
		panic(err)
	}
	fmt.Println("Migration completed")
}

func GetDB() *gorm.DB {
	return DB
}
