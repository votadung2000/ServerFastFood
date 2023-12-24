package ginCategory

import (
	"fastFood/common"
	bizCategory "fastFood/modules/category/biz"
	storageCategory "fastFood/modules/category/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindCategoryHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		store := storageCategory.NewSqlStorage(db)
		business := bizCategory.FindCategoryBiz(store)

		data, err := business.FindCategory(ctx.Request.Context(), id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
