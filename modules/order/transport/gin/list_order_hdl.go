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

func ListOrderHdl(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging

		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		paging.Process()

		var filter modelOrder.Filter

		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		store := storageOrder.NewSQLStorage(db)
		business := bizOrder.NewListOrderBiz(store)

		data, err := business.ListOrder(ctx.Request.Context(), &filter, &paging)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SuccessResponse(data, filter, paging))
	}
}
