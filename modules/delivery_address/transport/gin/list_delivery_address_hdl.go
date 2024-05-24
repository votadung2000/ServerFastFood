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

func ListDeliveryAddress(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging

		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		paging.Process()

		var filter modelDeliveryAddress.Filter

		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		user := ctx.MustGet(common.CurrentUser).(common.Requester)

		store := storageDeliveryAddress.NewSQLStorage(db)
		business := bizDeliveryAddress.NewListDeliveryAddressBiz(store)

		data, err := business.ListDeliveryAddress(ctx.Request.Context(), &filter, &paging, user.GetUserId())

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SuccessResponse(data, filter, paging))
	}
}
