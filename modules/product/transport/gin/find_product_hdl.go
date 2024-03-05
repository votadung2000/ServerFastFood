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

func FindProductHdl(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		store := storageProduct.NewSQLStorage(db)
		business := bizProduct.NewFindProductBiz(store)

		data, err := business.FindProduct(ctx.Request.Context(), id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
