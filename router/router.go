package router

import (
	category "example.com/m/controller/category"
	"example.com/m/database"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/category", category.CreateCategoryItem(database.Connections()))
		v1.GET("/category", category.GetAllCategoryItems(database.Connections()))
		v1.GET("/category/:id", category.GetDetailCategoryItem(database.Connections()))
		v1.PUT("/category/:id", category.UpdatesCategoryItem(database.Connections()))
		v1.DELETE("/category/:id", category.DeleteCategoryItem(database.Connections()))
	}

	router.Run()
}
