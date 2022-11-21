package category

import (
	"fmt"
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

		fmt.Println("Log", data.Select("Name").Create(&categoryItems))
		if err := data.Select("Name").Create(&categoryItems).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": categoryItems})
	}
}
