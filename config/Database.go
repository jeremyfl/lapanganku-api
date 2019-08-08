package config

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var err error

const (
	DRIVER = "mysql"
)

// Database connect
func Database() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open(DRIVER, ""+os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@/"+os.Getenv("DB_NAME")+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect to database")
	}

	return db
}
