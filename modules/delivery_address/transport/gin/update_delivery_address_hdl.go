package ginDeliveryAddress

import (
	"fastFood/common"
	bizDeliveryAddress "fastFood/modules/delivery_address/biz"
	modelDeliveryAddress "fastFood/modules/delivery_address/model"
	storageDeliveryAddress "fastFood/modules/delivery_address/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateDeliveryAddressHdl(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		user := ctx.MustGet(common.CurrentUser).(common.Requester)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		var data modelDeliveryAddress.DeliveryAddressUpdate

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		store := storageDeliveryAddress.NewSQLStorage(db)
		business := bizDeliveryAddress.NewUpdateDeliveryAddressBiz(store)

		if err := business.UpdateDeliveryAddress(ctx.Request.Context(), user.GetUserId(), id, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
