package ginDeliveryAddress

import (
	"fastFood/common"
	bizDeliveryAddress "fastFood/modules/delivery_address/biz"
	storageDeliveryAddress "fastFood/modules/delivery_address/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindDeliveryAddressDefaultHdl(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		store := storageDeliveryAddress.NewSQLStorage(db)
		business := bizDeliveryAddress.NewFindDeliveryAddressDefaultBiz(store)

		data, err := business.FindDeliveryAddressDefault(ctx.Request.Context())

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
