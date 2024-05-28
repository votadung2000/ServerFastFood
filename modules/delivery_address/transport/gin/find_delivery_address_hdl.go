package ginDeliveryAddress

import (
	"fastFood/common"
	bizDeliveryAddress "fastFood/modules/delivery_address/biz"
	storageDeliveryAddress "fastFood/modules/delivery_address/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindDeliveryAddressHdl(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet(common.CurrentUser).(common.Requester)

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		store := storageDeliveryAddress.NewSQLStorage(db)
		business := bizDeliveryAddress.NewFindDeliveryAddressBiz(store)

		data, err := business.FindDeliveryAddress(ctx.Request.Context(), id, user.GetUserId())

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
