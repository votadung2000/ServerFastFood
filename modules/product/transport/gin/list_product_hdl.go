package ginProduct

import (
	"fastFood/common"
	bizProduct "fastFood/modules/product/biz"
	modelProduct "fastFood/modules/product/model"
	storageProduct "fastFood/modules/product/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListProductHdl(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet(common.CurrentUser).(common.Requester)

		var queryString struct {
			common.Paging
			modelProduct.Filter
		}

		if err := ctx.ShouldBind(&queryString); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		queryString.Process()

		store := storageProduct.NewSQLStorage(db)
		business := bizProduct.NewListProductBiz(store)

		data, err := business.ListProduct(
			ctx.Request.Context(),
			user.GetUserId(),
			&queryString.Filter,
			&queryString.Paging,
		)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SuccessResponse(data, queryString.Filter, queryString.Paging))
	}
}
