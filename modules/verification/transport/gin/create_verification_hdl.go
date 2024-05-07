package ginVerification

import (
	"fastFood/common"
	"fastFood/components/tokenProvider"
	storageUser "fastFood/modules/user/storage"
	bizVerification "fastFood/modules/verification/biz"
	modelVerification "fastFood/modules/verification/model"
	storageVerification "fastFood/modules/verification/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateVerificationHdl(db *gorm.DB, tokenProvider tokenProvider.Provider) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data modelVerification.ParamsVerification

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		storeVerification := storageVerification.NewSQLStorage(db)
		storeUser := storageUser.NewSQLStorage(db)

		business := bizVerification.NewCreateVerificationBiz(storeUser, storeVerification, tokenProvider, 60*60*1)

		if err := business.CreateVerification(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
