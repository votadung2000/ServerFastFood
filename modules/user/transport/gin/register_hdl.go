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

func RegisterHdl(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data modelUser.UserCreate

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		store := storageUser.NewSQLStorage(db)
		md5 := common.NewMd5Hash()

		business := bizUser.NewCreateUserBiz(store, md5)

		if err := business.CreateUser(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
