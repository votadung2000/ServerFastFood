package ginCategory

import (
	"fastFood/common"
	bizCategory "fastFood/modules/category/biz"
	modelCategory "fastFood/modules/category/model"
	storageCategory "fastFood/modules/category/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListCategoryHdl(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging

		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		paging.Process()

		var filter modelCategory.Filter

		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		store := storageCategory.NewSqlStorage(db)
		business := bizCategory.NewListCategoryBiz(store)

		data, err := business.ListCategory(ctx.Request.Context(), &filter, &paging)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SuccessResponse(data, filter, paging))
	}
}
