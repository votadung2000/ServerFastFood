package router

import (
	"log"
	"os"

	category "example.com/m/controller/category"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Router() {
	env := os.Getenv("API_HOST")

	dsn := "root:pass-server-mysql@tcp(" + env + ":3306)/go_fast_food_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	log.Println("Connected to MySQL:", db)

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/category", category.CreateCategoryItem(db))
		v1.GET("/category", category.GetAllCategoryItem(db))
	}

	router.Run()
}
