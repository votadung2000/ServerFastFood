package ginUpload

import (
	"fastFood/common"
	bizUpload "fastFood/modules/upload/biz"
	storageUpload "fastFood/modules/upload/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindImageHdl(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		store := storageUpload.NewSQLStorage(db)
		business := bizUpload.NewFindImageBiz(store)

		data, err := business.FindImage(ctx.Request.Context(), id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
