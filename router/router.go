package router

import (
	favorite "fastFood/controller/favorite"
	product "fastFood/controller/product"
	user "fastFood/controller/user"
	"fastFood/database"
	ginCategory "fastFood/modules/category/transport/gin"
	ginProduct "fastFood/modules/product/transport/gin"

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
		v1.POST("/category", ginCategory.CreateCategoryHandler(db))
		v1.GET("/category", ginCategory.ListCategoryHandler(db))
		v1.GET("/category/:id", ginCategory.FindCategoryHandler(db))
		v1.PUT("/category/:id", ginCategory.UpdateCategoryHandler(db))
		v1.DELETE("/category/:id", ginCategory.DeleteCategoryHandler(db))
	}
	{
		v1.POST("/product", ginProduct.CreateProductHandler(db))
		v1.GET("/product", product.GetAllProductItems(db))
		v1.GET("/product/:id", ginProduct.FindProductHandler(db))
		v1.PUT("/product/:id", ginProduct.UpdateProductHandler(db))
		v1.DELETE("/product/:id", product.DeleteProductItem(db))
	}
	{
		v1.POST("/favorite", favorite.CreateFavorite(db))
		v1.GET("/favorite", favorite.GetAllFavorites(db))
		v1.POST("/favorite/:id", favorite.DeleteFavorite(db))
	}

	router.Run()
}
