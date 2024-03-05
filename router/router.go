package router

import (
	"fastFood/common"
	jwtProvider "fastFood/components/tokenProvider/jwt"
	"fastFood/database"
	"fastFood/middleware"
	ginCategory "fastFood/modules/category/transport/gin"
	ginFavorite "fastFood/modules/favorite/transport/gin"
	ginProduct "fastFood/modules/product/transport/gin"
	storageUser "fastFood/modules/user/storage"
	ginUser "fastFood/modules/user/transport/gin"
	"os"

	"github.com/gin-gonic/gin"
)

func Router() {
	db := database.Connections()

	gin.SetMode(gin.ReleaseMode)

	secret := os.Getenv("SECRET_JWT")
	tokenProvider := jwtProvider.NewJwtProvider(common.JWT, secret)
	authStore := storageUser.NewSQLStorage(db)
	middlewareAuth := middleware.RequireAuth(authStore, tokenProvider)

	router := gin.Default()
	router.Use(middleware.Recover())
	router.Static("/static", "./static")

	v1 := router.Group("/v1")
	{
		// upload := v1.Group("/upload", middlewareAuth)
		// {
		// 	upload.PUT("", upload.Upload(db))
		// }

		{
			v1.POST("/login", ginUser.LoginHdl(db, tokenProvider))
			v1.POST("/register", ginUser.RegisterHdl(db))
			v1.GET("/profile", middlewareAuth, ginUser.ProfileUserHandler(db))
		}

		category := v1.Group("/category", middlewareAuth)
		{
			category.POST("", ginCategory.CreateCategoryHdl(db))
			category.GET("", ginCategory.ListCategoryHandler(db))
			category.GET("/:id", ginCategory.FindCategoryHandler(db))
			category.PATCH("/:id", ginCategory.UpdateCategoryHandler(db))
			category.DELETE("/:id", ginCategory.DeleteCategoryHandler(db))
		}

		product := v1.Group("/product", middlewareAuth)
		{
			product.POST("", ginProduct.CreateProductHandler(db))
			product.GET("", ginProduct.ListProductHandler(db))
			product.GET("/:id", ginProduct.FindProductHandler(db))
			product.PATCH("/:id", ginProduct.UpdateProductHandler(db))
			product.DELETE("/:id", ginProduct.DeleteProductHandler(db))
		}

		favorite := v1.Group("/favorite", middlewareAuth)
		{
			favorite.POST("", ginFavorite.CreateFavoriteHandler(db))
			// v1.GET("/favorite", favorite.GetAllFavorites(db))
			favorite.DELETE("/:id", ginFavorite.DeleteFavoriteHandler(db))
		}
	}

	router.Run(":3000")
}
