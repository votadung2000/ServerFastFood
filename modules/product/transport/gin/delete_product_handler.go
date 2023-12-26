package ginProduct

import (
	"fastFood/common"
	bizProduct "fastFood/modules/product/biz"
	storageProduct "fastFood/modules/product/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteProductHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		store := storageProduct.NewSQLStorage(db)
		business := bizProduct.DeleteProductBiz(store)

		if err := business.DeleteProduct(ctx.Request.Context(), id); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}