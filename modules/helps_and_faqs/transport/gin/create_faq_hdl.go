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

func CreateFAQHdl(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data modelFAQ.FAQCreate

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		store := storageFAQ.NewSQLStorage(db)
		business := bizFAQ.NewCreateFAQBiz(store)

		if err := business.CreateFAQ(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
