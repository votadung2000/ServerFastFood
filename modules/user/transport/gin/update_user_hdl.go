package ginUser

import (
	"fastFood/common"
	bizUser "fastFood/modules/user/biz"
	modelUser "fastFood/modules/user/model"
	storageUser "fastFood/modules/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateUserHdl(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet(common.CurrentUser).(common.Requester)

		var data modelUser.UserUpdate

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		store := storageUser.NewSQLStorage(db)
		business := bizUser.NewUpdateUserBiz(store)

		if err := business.UpdateUser(ctx.Request.Context(), user.GetUserId(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
