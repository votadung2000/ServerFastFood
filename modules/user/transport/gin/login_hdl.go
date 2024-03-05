package ginUser

import (
	"fastFood/common"
	"fastFood/components/tokenProvider"
	bizUser "fastFood/modules/user/biz"
	modelUser "fastFood/modules/user/model"
	storageUser "fastFood/modules/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoginHdl(db *gorm.DB, tokenProvider tokenProvider.Provider) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data modelUser.Login

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		store := storageUser.NewSQLStorage(db)
		md5 := common.NewMd5Hash()

		business := bizUser.NewLoginBiz(store, md5, tokenProvider, 60*60*24*7)

		account, err := business.Login(ctx.Request.Context(), &data)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
