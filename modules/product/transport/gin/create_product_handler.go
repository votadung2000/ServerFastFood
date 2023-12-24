package ginProduct

import (
	"fastFood/common"
	bizProduct "fastFood/modules/product/biz"
	modelProduct "fastFood/modules/product/model"
	storageProduct "fastFood/modules/product/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProductHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data modelProduct.ProductCreate

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		store := storageProduct.NewSQLStorage(db)
		business := bizProduct.CreateProductBiz(store)

		if err := business.CreateProduct(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
