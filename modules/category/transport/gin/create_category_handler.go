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

func CreateCategoryHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data modelCategory.CategoryCreate

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		store := storageCategory.NewSqlStorage(db)
		business := bizCategory.CreateCategoryBiz(store)

		if err := business.CreateCategory(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
