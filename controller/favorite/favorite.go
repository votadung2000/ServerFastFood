package favorite

import (
	"net/http"
	"strconv"

	"fastFood/models"

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

		var resultFavorites []models.Favorite

		proGetFavorites := data.Table(models.Favorite{}.TableFavorites())

		if filterF.UserId > 0 {
			proGetFavorites.Where("user_id = ?", filterF.UserId)
		}

		// proGetFavorites.Count(&filterF.Total)
		proGetFavorites.Limit(filterF.Limit)
		proGetFavorites.Offset(offset)
		proGetFavorites.Order("id desc")
		proGetFavorites.Find(&resultFavorites)

		if err := proGetFavorites.Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		ids := []int64{}

		for _, item := range resultFavorites {
			ids = append(ids, int64(item.Id))
		}

		var resultProducts []models.Product

		proGetProducts := data.Table(models.Product{}.TableProducts())
		proGetProducts.Where("status = ?", 1)

		if int64(len(ids)) > 0 {
			proGetProducts.Where("id IN ?", ids)
		}

		if filterF.CategoryId > 0 {
			proGetProducts.Where("category_id = ?", filterF.CategoryId)
		}

		proGetFavorites.Count(&filterF.Total)
		proGetProducts.Limit(filterF.Limit)
		proGetProducts.Offset(offset)
		proGetProducts.Order("id desc")
		proGetProducts.Find(&resultProducts)

		if err := proGetProducts.Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		responseClient.Total = filterF.Total
		responseClient.Data = resultProducts

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
