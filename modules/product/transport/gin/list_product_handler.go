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

func ListProductHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging

		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		paging.Process()

		var filter modelProduct.Filter

		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		store := storageProduct.NewSQLStorage(db)
		business := bizProduct.ListProductBiz(store)

		data, err := business.ListProduct(
			ctx.Request.Context(),
			&filter,
			&paging,
		)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, common.SuccessResponse(data, filter, paging))
	}
}
