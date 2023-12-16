package category

import (
	"net/http"
	"strconv"

	"fastFood/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProduct(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var productItem models.Product

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
		var filterP models.FormatGetListProducts
		var responseClient models.FormatResponse

		if err := context.ShouldBind(&filterP); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		if filterP.Page <= 0 {
			filterP.Page = 1
		}

		if filterP.Limit <= 0 {
			filterP.Limit = 10
		}

		offset := (filterP.Page - 1) * filterP.Limit

		var results []models.Product

		// ids := []int64{1, 2, 3, 4, 5}

		resultGetAll := data.Table(models.Product{}.TableProducts())
		resultGetAll.Where("status = ?", 1)
		// resultGetAll.Where("id IN ?", ids)

		if filterP.CategoryId > 0 {
			resultGetAll.Where("category_id = ?", filterP.CategoryId)
		}

		if filterP.Name != "" {
			resultGetAll.Where("name LIKE ?", "%"+filterP.Name+"%")
		}

		resultGetAll.Count(&filterP.Total)
		resultGetAll.Limit(filterP.Limit)
		resultGetAll.Offset(offset)
		resultGetAll.Order("id desc")
		resultGetAll.Find(&results)

		if err := resultGetAll.Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		responseClient.Total = filterP.Total
		responseClient.Data = results

		context.JSON(http.StatusOK, responseClient)
	}
}

func GetDetailProductItem(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var productItem models.Product

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

		var productItem models.Product

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

		if err := data.Table(models.Product{}.TableProducts()).
			Where("id = ?", id).
			Delete(nil).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
