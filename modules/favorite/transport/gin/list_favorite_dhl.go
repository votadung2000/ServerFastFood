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

func ListFavoriteHdl(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet(common.CurrentUser).(common.Requester)

		var paging common.Paging

		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		paging.Process()

		var filter modelFavorite.Filter

		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		store := storageFavorite.NewSQLStorage(db)
		business := bizFavorite.NewListFavoriteBiz(store)

		data, err := business.ListFavorite(ctx.Request.Context(), user.GetUserId(), &filter, &paging)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SuccessResponse(data, filter, paging))
	}
}
