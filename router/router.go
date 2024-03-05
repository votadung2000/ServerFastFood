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

	// "fastFood/middleware"
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

		{
			v1.POST("/category", ginCategory.CreateCategoryHandler(db))
			v1.GET("/category", ginCategory.ListCategoryHandler(db))
			v1.GET("/category/:id", ginCategory.FindCategoryHandler(db))
			v1.PATCH("/category/:id", ginCategory.UpdateCategoryHandler(db))
			v1.DELETE("/category/:id", ginCategory.DeleteCategoryHandler(db))
		}
		{
			v1.POST("/product", ginProduct.CreateProductHandler(db))
			v1.GET("/product", ginProduct.ListProductHandler(db))
			v1.GET("/product/:id", ginProduct.FindProductHandler(db))
			v1.PATCH("/product/:id", ginProduct.UpdateProductHandler(db))
			v1.DELETE("/product/:id", ginProduct.DeleteProductHandler(db))
		}
		{
			v1.POST("/favorite", ginFavorite.CreateFavoriteHandler(db))
			// v1.GET("/favorite", favorite.GetAllFavorites(db))
			v1.DELETE("/favorite/:id", ginFavorite.DeleteFavoriteHandler(db))
		}
	}

	router.Run(":3000")
}
