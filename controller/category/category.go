package category

import (
	"net/http"
	"strconv"
	"strings"

	"example.com/m/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCategoryItem(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var categoryItem models.Categories

		if err := context.ShouldBind(&categoryItem); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		// Name
		categoryItem.Name = strings.TrimSpace(categoryItem.Name)

		if categoryItem.Name == "" {
			context.JSON(http.StatusBadRequest, gin.H{"Error": "Name cannot be blank"})
			return
		}

		if err := data.Select("Name").Create(&categoryItem).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": categoryItem})
	}
}

func GetAllCategoryItems(data *gorm.DB) gin.HandlerFunc {
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

		offset := (paging.Page - 1) * paging.Limit

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

func GetDetailCategoryItem(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var categoryItem models.Categories

		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		if err := data.Where("id = ?", id).First(&categoryItem).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": categoryItem})
	}
}

func UpdatesCategoryItem(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		var categoryItem models.Categories

		if err := context.ShouldBind(&categoryItem); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		if err := data.Where("id = ?", id).Updates(&categoryItem).First(&categoryItem).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": categoryItem})
	}
}

func DeleteCategoryItem(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		if err := data.Table(models.Categories{}.TableCategory()).
			Where("id = ?", id).
			Delete(nil).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
