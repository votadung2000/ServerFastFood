package router

import (
	category "example.com/m/controller/category"
	"example.com/m/controller/user"
	"example.com/m/database"
	"github.com/gin-gonic/gin"
)

func Router() {
	db := database.Connections()

	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/user/login", user.HandleLogin(db))
		v1.POST("/user/register", user.HandleRegister(db))
		v1.GET("/user/:id", user.GetDetailUserItem(db))
		v1.PUT("/user/:id", user.UpdatesUserItem(db))
		v1.DELETE("/user/:id", user.DeleteUserItem(db))
	}
	{
		v1.POST("/category", category.CreateCategoryItem(db))
		v1.GET("/category", category.GetAllCategoryItems(db))
		v1.GET("/category/:id", category.GetDetailCategoryItem(db))
		v1.PUT("/category/:id", category.UpdatesCategoryItem(db))
		v1.DELETE("/category/:id", category.DeleteCategoryItem(db))
	}

	router.Run()
}
