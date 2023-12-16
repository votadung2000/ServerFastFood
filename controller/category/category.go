package category

import (
	"net/http"
	"strconv"

	"fastFood/components"
	"fastFood/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCategoryItem(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var categoryItem models.Categories

		if err := context.ShouldBind(&categoryItem); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		// Name
		categoryItem.Name = components.Sanitize(categoryItem.Name)

		if categoryItem.Name == "" {
			context.JSON(http.StatusBadRequest, gin.H{"Message": "Name cannot be blank"})
			return
		}

		if err := data.Select("Name").Create(&categoryItem).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": categoryItem})
	}
}

func GetAllCategoryItems(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var filterC models.FormatGetList
		var responseClient models.FormatResponse

		if err := context.ShouldBind(&filterC); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		if filterC.Page <= 0 {
			filterC.Page = 1
		}

		if filterC.Limit <= 0 {
			filterC.Limit = 10
		}

		offset := (filterC.Page - 1) * filterC.Limit

		var result []models.Categories

		if err := data.Table(models.Categories{}.TableCategory()).
			Count(&filterC.Total).
			Limit(filterC.Limit).
			Offset(offset).
			Find(&result).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		responseClient.Total = filterC.Total
		responseClient.Data = result

		context.JSON(http.StatusOK, responseClient)
	}
}

func GetDetailCategoryItem(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var categoryItem models.Categories

		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		if err := data.Where("id = ?", id).First(&categoryItem).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": categoryItem})
	}
}

func UpdatesCategoryItem(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		var categoryItem models.Categories

		if err := context.ShouldBind(&categoryItem); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		if err := data.Where("id = ?", id).Updates(&categoryItem).First(&categoryItem).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": categoryItem})
	}
}

func DeleteCategoryItem(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		if err := data.Table(models.Categories{}.TableCategory()).
			Where("id = ?", id).
			Delete(nil).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
