package ginDeliveryAddress

import (
	"fastFood/common"
	bizDeliveryAddress "fastFood/modules/delivery_address/biz"
	modelDeliveryAddress "fastFood/modules/delivery_address/model"
	storageDeliveryAddress "fastFood/modules/delivery_address/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateDeliveryAddressHdl(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet(common.CurrentUser).(common.Requester)

		var data modelDeliveryAddress.CreateDeliveryAddress

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		store := storageDeliveryAddress.NewSQLStorage(db)
		business := bizDeliveryAddress.NewCreateDeliveryAddressBiz(store)

		if err := business.CreateDeliveryAddress(ctx.Request.Context(), user.GetUserId(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
