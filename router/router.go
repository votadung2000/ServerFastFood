package router

import (
	category "fastFood/controller/category"
	favorite "fastFood/controller/favorite"
	product "fastFood/controller/product"
	user "fastFood/controller/user"
	"fastFood/database"
	ginCategory "fastFood/modules/category/transport/gin"

	// "fastFood/middleware"
	"github.com/gin-gonic/gin"
)

func Router() {
	db := database.Connections()

	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// router.Use(middleware.Authentication())

	v1 := router.Group("/v1")
	{
		v1.POST("/login", user.HandleLogin(db))
		v1.POST("/register", user.HandleRegister(db))
		v1.GET("/user/:id", user.GetDetailUserItem(db))
		v1.PUT("/user/:id", user.UpdatesUserItem(db))
		v1.DELETE("/user/:id", user.DeleteUserItem(db))
	}

	{
		v1.POST("/category", category.CreateCategoryItem(db))
		v1.GET("/category", ginCategory.ListCategoryHandler(db))
		v1.GET("/category/:id", ginCategory.FindCategoryHandler(db))
		v1.PUT("/category/:id", category.UpdatesCategoryItem(db))
		v1.DELETE("/category/:id", category.DeleteCategoryItem(db))
	}
	{
		v1.POST("/product", product.CreateProduct(db))
		v1.GET("/product", product.GetAllProductItems(db))
		v1.GET("/product/:id", product.GetDetailProductItem(db))
		v1.PUT("/product/:id", product.UpdatesProductItem(db))
		v1.DELETE("/product/:id", product.DeleteProductItem(db))
	}
	{
		v1.POST("/favorite", favorite.CreateFavorite(db))
		v1.GET("/favorite", favorite.GetAllFavorites(db))
		v1.POST("/favorite/:id", favorite.DeleteFavorite(db))
	}

	router.Run()
}
