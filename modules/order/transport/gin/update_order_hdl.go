package ginOrder

import (
	"fastFood/common"
	bizOrder "fastFood/modules/order/biz"
	modelOrder "fastFood/modules/order/model"
	storageOrder "fastFood/modules/order/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateOrderHdl(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		var data modelOrder.UpdateOrder

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		store := storageOrder.NewSQLStorage(db)
		business := bizOrder.NewUpdateOrderBiz(store)

		if err := business.UpdateOrder(ctx.Request.Context(), id, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
