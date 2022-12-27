package category

import (
	"net/http"
	"strconv"

	"example.com/m/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProduct(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var productItem models.Products

		if err := context.ShouldBind(&productItem); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		if productItem.Name == "" {
			context.JSON(http.StatusBadRequest, gin.H{"Message": "Name cannot be blank"})
			return
		}

		if err := data.Select("Name").Create(&productItem).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": productItem})
	}
}

func GetAllProductItems(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var paging models.FormatGetList

		if err := context.ShouldBind(&paging); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		if paging.Page <= 0 {
			paging.Page = 1
		}

		if paging.Limit <= 0 {
			paging.Limit = 10
		}

		offset := (paging.Page - 1) * paging.Limit

		var result []models.Products

		if err := data.Table(models.Products{}.TableProducts()).
			Count(&paging.Total).
			Limit(paging.Limit).
			Offset(offset).
			Order("id desc").
			Find(&result).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": result})
	}
}

func GetDetailProductItem(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var productItem models.Products

		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		if err := data.Where("id = ?", id).First(&productItem).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": productItem})
	}
}

func UpdatesProductItem(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		var productItem models.Products

		if err := context.ShouldBind(&productItem); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		if err := data.Where("id = ?", id).Updates(&productItem).First(&productItem).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": productItem})
	}
}

func DeleteProductItem(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		if err := data.Table(models.Products{}.TableProducts()).
			Where("id = ?", id).
			Delete(nil).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
