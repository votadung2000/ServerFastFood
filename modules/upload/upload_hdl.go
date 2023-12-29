package upload

import (
	"fastFood/common"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Upload(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fileHeader, err := ctx.FormFile("file")

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		dst := fmt.Sprintf("./static/%d.%s", time.Now().UTC().UnixNano(), fileHeader.Filename)

		if err := ctx.SaveUploadedFile(fileHeader, dst); err != nil {
			return
		}

		img := common.Image{
			Id:        1,
			Url:       dst,
			Width:     1200,
			Height:    1200,
			CloudName: "local",
			Extension: "",
		}

		img.Fulfil("http://localhost:3000")

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
