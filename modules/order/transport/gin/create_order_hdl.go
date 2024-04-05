package ginOrder

import (
	"fastFood/common"
	bizOrder "fastFood/modules/order/biz"
	modelOrder "fastFood/modules/order/model"
	storageOrder "fastFood/modules/order/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrderHdl(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data modelOrder.CreateOrder

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		store := storageOrder.NewSQLStorage(db)
		business := bizOrder.NewCreateOrder(store)

		if err := business.CreateOrder(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
