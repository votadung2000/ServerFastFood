package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connections() *gorm.DB {
	env := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	charset := os.Getenv("DB_CHARSET")
	parseTime := os.Getenv("DB_PARSE_TIME")
	loc := os.Getenv("DB_LOC")

	dsn := user + ":" + pass + "@tcp(" + env + ":" + port + ")/" + name + "?charset=" + charset + "&parseTime=" + parseTime + "&loc=" + loc

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	log.Println("Connected To MySQL Success")
	return db
}
