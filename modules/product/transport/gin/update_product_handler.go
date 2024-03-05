package ginProduct

import (
	"fastFood/common"
	bizProduct "fastFood/modules/product/biz"
	modelProduct "fastFood/modules/product/model"
	storageProduct "fastFood/modules/product/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateProductHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		var data modelProduct.ProductUpdate

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		store := storageProduct.NewSQLStorage(db)
		business := bizProduct.UpdateProductBiz(store)

		if err := business.UpdateProduct(ctx.Request.Context(), id, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
