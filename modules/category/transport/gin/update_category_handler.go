package ginCategory

import (
	"fastFood/common"
	bizCategory "fastFood/modules/category/biz"
	modelCategory "fastFood/modules/category/model"
	storageCategory "fastFood/modules/category/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateCategoryHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		var dataUpdate modelCategory.CategoryUpdate

		if err := ctx.ShouldBind(&dataUpdate); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		store := storageCategory.NewSqlStorage(db)
		business := bizCategory.UpdateCategoryBiz(store)

		if err := business.UpdateCategory(ctx.Request.Context(), id, &dataUpdate); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
