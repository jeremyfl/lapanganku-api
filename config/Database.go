package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var err error

// Database connect
func Database() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open("mysql", ""+os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@/"+os.Getenv("DB_NAME")+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("Can't connect to database", err)
	}

	return db
}
