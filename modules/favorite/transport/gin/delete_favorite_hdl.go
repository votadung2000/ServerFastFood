package ginFavorite

import (
	"fastFood/common"
	bizFavorite "fastFood/modules/favorite/biz"
	storageFavorite "fastFood/modules/favorite/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteFavoriteHdl(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		store := storageFavorite.NewSQLStorage(db)
		business := bizFavorite.DeleteFavoriteBiz(store)

		if err := business.DeleteFavorite(ctx.Request.Context(), id); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
