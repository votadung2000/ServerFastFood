package router

import (
	"fastFood/common"
	jwtProvider "fastFood/components/tokenProvider/jwt"
	"fastFood/database"
	"fastFood/middleware"
	ginCategory "fastFood/modules/category/transport/gin"
	ginDeliveryAddress "fastFood/modules/delivery_address/transport/gin"
	ginFavorite "fastFood/modules/favorite/transport/gin"
	ginFAQ "fastFood/modules/helps_and_faqs/transport/gin"
	ginOrder "fastFood/modules/order/transport/gin"
	ginProduct "fastFood/modules/product/transport/gin"
	ginUpload "fastFood/modules/upload/transport/gin"
	storageUser "fastFood/modules/user/storage"
	ginUser "fastFood/modules/user/transport/gin"
	ginVerification "fastFood/modules/verification/transport/gin"
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
		{
			v1.POST("/login", ginUser.LoginHdl(db, tokenProvider))
			v1.POST("/register", ginUser.RegisterHdl(db))
			v1.GET("/profile", middlewareAuth, ginUser.ProfileUserHandler(db))
			v1.PATCH("/update_user", middlewareAuth, ginUser.UpdateUserHdl(db))
			v1.PATCH("/update_password", ginUser.UpdatePasswordHdl(db))
		}

		category := v1.Group("/category", middlewareAuth)
		{
			category.POST("", ginCategory.CreateCategoryHdl(db))
			category.GET("", ginCategory.ListCategoryHdl(db))
			category.GET("/:id", ginCategory.FindCategoryHdl(db))
			category.PATCH("/:id", ginCategory.UpdateCategoryHdl(db))
			category.DELETE("/:id", ginCategory.DeleteCategoryHdl(db))
		}

		deliveryAddress := v1.Group("/delivery_address", middlewareAuth)
		{
			deliveryAddress.POST("", ginDeliveryAddress.CreateDeliveryAddressHdl(db))
			deliveryAddress.GET("", ginDeliveryAddress.ListDeliveryAddress(db))
			deliveryAddress.GET("/default", ginDeliveryAddress.FindDeliveryAddressDefaultHdl(db))
			deliveryAddress.GET("/:id", ginDeliveryAddress.FindDeliveryAddressHdl(db))
			deliveryAddress.PATCH("/:id", ginDeliveryAddress.UpdateDeliveryAddressHdl(db))
			deliveryAddress.DELETE("/:id", ginDeliveryAddress.DeleteDeliveryAddressHdl(db))
		}

		product := v1.Group("/product", middlewareAuth)
		{
			product.POST("", ginProduct.CreateProductHdl(db))
			product.GET("", ginProduct.ListProductHdl(db))
			product.GET("/:id", ginProduct.FindProductHdl(db))
			product.PATCH("/:id", ginProduct.UpdateProductHdl(db))
			product.DELETE("/:id", ginProduct.DeleteProductHdl(db))
		}

		favorite := v1.Group("/favorite", middlewareAuth)
		{
			favorite.POST("", ginFavorite.CreateFavoriteHdl(db))
			favorite.POST("/cd_favorite", ginFavorite.CDFavoriteHdl(db))
			favorite.GET("", ginFavorite.ListFavoriteHdl(db))
			favorite.DELETE("/:id", ginFavorite.DeleteFavoriteHdl(db))
		}

		order := v1.Group("/order", middlewareAuth)
		{
			order.POST("", ginOrder.CreateOrderHdl(db))
			order.GET("", ginOrder.ListOrderHdl(db))
			order.GET("/:id", ginOrder.FindOrderHdl(db))
			order.PATCH("/:id", ginOrder.UpdateOrderHdl(db))
		}

		upload := v1.Group("/upload", middlewareAuth)
		{
			upload.POST("", ginUpload.CreateImageHdl(db))
			upload.GET("/:id", ginUpload.FindImageHdl(db))
		}

		verification := v1.Group("/verification")
		{
			verification.POST("", ginVerification.CreateVerificationHdl(db, tokenProvider))
		}

		faq := v1.Group("/helps_and_faqs", middlewareAuth)
		{
			faq.POST("", ginFAQ.CreateFAQHdl(db))
			faq.GET("", ginFAQ.ListFAQHdl(db))
			faq.DELETE("/:id", ginFAQ.DeleteFAQHdl(db))
		}
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
