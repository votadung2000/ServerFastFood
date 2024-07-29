package ginFAQ

import (
	"fastFood/common"
	bizFAQ "fastFood/modules/helps_and_faqs/biz"
	modelFAQ "fastFood/modules/helps_and_faqs/model"
	storageFAQ "fastFood/modules/helps_and_faqs/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListFAQHdl(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var queryString struct {
			common.Paging
			modelFAQ.Filter
		}

		if err := ctx.ShouldBind(&queryString); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		queryString.Process()

		store := storageFAQ.NewSQLStorage(db)
		business := bizFAQ.NewListFAQBiz(store)

		data, err := business.ListFAQ(
			ctx.Request.Context(),
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
