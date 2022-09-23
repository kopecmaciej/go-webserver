package lib

import (
	"fmt"
	"go-web-server/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dns = &config.GlobalConfig.Database.Url
	DB  *gorm.DB
)

func Open() *gorm.DB {
	db, err := gorm.Open(postgres.Open(*dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error while opening connection to DB: %v", err)
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
