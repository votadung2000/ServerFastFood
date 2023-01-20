package favorite

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/m/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateFavorite(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var favoriteItem models.Favorite

		if err := context.ShouldBind(&favoriteItem); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		if favoriteItem.UserId <= 0 || favoriteItem.ProductId <= 0 {
			context.JSON(http.StatusBadRequest, gin.H{"Message": "Cannot be blank"})
			return
		}

		if err := data.Select("UserId", "ProductId").Create(&favoriteItem).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}

func GetAllFavorites(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var filterF models.FormatGetFavorites
		var responseClient models.FormatResponse

		if err := context.ShouldBind(&filterF); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		if filterF.Page <= 0 {
			filterF.Page = 1
		}

		if filterF.Limit <= 0 {
			filterF.Limit = 10
		}

		offset := (filterF.Page - 1) * filterF.Limit

		var resultsFavorites []models.Favorite

		if err := data.Table(models.Favorite{}.TableFavorites()).
			Where("user_id = ?", filterF.UserId).
			Count(&filterF.Total).
			Limit(filterF.Limit).
			Offset(offset).
			Order("id desc").
			Find(&resultsFavorites).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		fmt.Println("resultsFavorites", resultsFavorites)

		responseClient.Total = filterF.Total
		responseClient.Data = resultsFavorites

		context.JSON(http.StatusOK, responseClient)
	}
}

func DeleteFavorite(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		if err := data.Table(models.Favorite{}.TableFavorites()).
			Where("id = ?", id).
			Delete(nil).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
