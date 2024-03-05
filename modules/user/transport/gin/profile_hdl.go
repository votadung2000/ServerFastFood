package ginUser

import (
	"fastFood/common"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProfileUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet(common.CurrentUser)

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(user))
	}
}
