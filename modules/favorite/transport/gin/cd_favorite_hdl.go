package ginFavorite

import (
	"fastFood/common"
	bizFavorite "fastFood/modules/favorite/biz"
	modelFavorite "fastFood/modules/favorite/model"
	storageFavorite "fastFood/modules/favorite/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CDFavoriteHdl(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data modelFavorite.FavoriteCreate

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		store := storageFavorite.NewSQLStorage(db)
		business := bizFavorite.CDFavoriteBiz(store)

		if err := business.CDFavorite(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
