package category

import (
	"net/http"
	"strings"

	"example.com/m/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCategoryItem(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var categoryItems models.Categories

		if err := context.ShouldBind(&categoryItems); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		// Name
		categoryItems.Name = strings.TrimSpace(categoryItems.Name)

		if categoryItems.Name == "" {
			context.JSON(http.StatusBadRequest, gin.H{"Error": "Name cannot be blank"})
			return
		}

		if err := data.Select("Name").Create(&categoryItems).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": categoryItems})
	}
}

func GetAllCategoryItem(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var paging models.FormatGetList

		if err := context.ShouldBind(&paging); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		if paging.Page <= 0 {
			paging.Page = 1
		}

		if paging.Limit <= 0 {
			paging.Limit = 10
		}

		offset := paging.Page * paging.Limit

		var result []models.Categories

		if err := data.Table(models.Categories{}.TableCategory()).
			Count(&paging.Total).
			Limit(paging.Limit).
			Offset(offset).
			Order("id desc").
			Find(&result).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": result})
	}
}
