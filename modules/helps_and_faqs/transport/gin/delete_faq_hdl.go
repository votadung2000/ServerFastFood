package ginFAQ

import (
	"fastFood/common"
	bizFAQ "fastFood/modules/helps_and_faqs/biz"
	storageFAQ "fastFood/modules/helps_and_faqs/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteFAQHdl(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		store := storageFAQ.NewSQLStorage(db)
		business := bizFAQ.NewDeleteFAQBiz(store)

		if err := business.DeleteFAQ(ctx.Request.Context(), id); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
