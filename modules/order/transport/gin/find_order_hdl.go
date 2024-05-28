package ginOrder

import (
	"fastFood/common"
	bizOrder "fastFood/modules/order/biz"
	storageOrder "fastFood/modules/order/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindOrderHdl(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet(common.CurrentUser).(common.Requester)

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		store := storageOrder.NewSQLStorage(db)
		business := bizOrder.NewFindOrderBiz(store)

		data, err := business.FindOrder(ctx.Request.Context(), id, user.GetUserId())

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
