package ginOrder

import (
	"fastFood/common"
	modelOrder "fastFood/modules/order/model"
	repoOrder "fastFood/modules/order/repository"
	storageOrder "fastFood/modules/order/storage"
	storageOrderItem "fastFood/modules/order_item/storage"
	storageProduct "fastFood/modules/product/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrderHdl(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data modelOrder.OrderParams

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		storeOrder := storageOrder.NewSQLStorage(db)
		storeOrderItem := storageOrderItem.NewSQLStorage(db)
		storeProduct := storageProduct.NewSQLStorage(db)

		repo := repoOrder.NewCreateOrderRepo(storeOrder, storeOrderItem, storeProduct)

		if err := repo.CreateOrder(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
