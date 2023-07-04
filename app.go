package config

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_STR")
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to Connect")
	}
	DB = d
}

func GetDB() *gorm.DB {
	return DB
}
