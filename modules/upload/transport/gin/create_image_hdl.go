package ginUpload

import (
	"fastFood/common"
	bizUpload "fastFood/modules/upload/biz"
	modelUpload "fastFood/modules/upload/model"
	storageUpload "fastFood/modules/upload/storage"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateImageHdl(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var img common.Image

		fileHeader, err := ctx.FormFile("file")

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		file, err := fileHeader.Open()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		typePathStr := ctx.PostForm("typePath")

		// if typePathStr == "" {
		// 	ctx.JSON(http.StatusBadRequest, modelUpload.ErrMissingField(err))
		// 	return
		// }

		typePathInt := common.TYPE_IMG_OTHER

		if typePathStr != "" {
			typePathInt, err = strconv.Atoi(typePathStr)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
				return
			}
		}

		dstLocal := generateFilePath(typePathInt, fileHeader.Filename)

		if err := ctx.SaveUploadedFile(fileHeader, dstLocal); err != nil {
			ctx.JSON(http.StatusBadRequest, modelUpload.ErrCannotSaveFile(err))
			return
		}

		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInternalRequest(err))
			return
		}

		store := storageUpload.NewSQLStorage(db)
		business := bizUpload.NewCreateImageBiz(store)

		if err := business.CreateImage(ctx.Request.Context(), dataBytes, fileHeader.Filename, dstLocal, &img); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}

func generateFilePath(typePathInt int, filename string) string {
	switch typePathInt {
	case common.TYPE_IMG_PROFILE:
		return filepath.Join("static", "profile", filename)
	case common.TYPE_IMG_PRODUCT:
		return filepath.Join("static", "product", filename)
	case common.TYPE_IMG_CATEGORY:
		return filepath.Join("static", "category", filename)
	default:
		// Mặc định, nếu không có key phù hợp, sẽ lưu vào thư mục static với tên ngẫu nhiên
		return filepath.Join("static", fmt.Sprintf("%d_%s", time.Now().UTC().UnixNano(), filename))
	}
}
