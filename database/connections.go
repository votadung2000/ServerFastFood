package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connections() *gorm.DB {
	env := os.Getenv("API_HOST")

	dsn := "root:pass-server-mysql@tcp(" + env + ":3306)/go_fast_food_db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	log.Println("Connected To MySQL Success")
	return db
}
