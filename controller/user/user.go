package user

import (
	"net/http"
	"strconv"

	"example.com/m/components"
	"example.com/m/models"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var userItem models.Users

		if err := context.ShouldBind(&userItem); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		if govalidator.IsNull(userItem.UserName) || govalidator.IsNull(userItem.PassWord) {
			context.JSON(http.StatusBadRequest, gin.H{"Message": "Data Can Not Empty"})
			return
		}

		userItem.UserName = components.Sanitize(userItem.UserName)
		userItem.PassWord = components.Sanitize(userItem.PassWord)

		var otherUserItem []models.Users

		if err := data.Table(models.Users{}.TableUsers()).
			Where("username = ?", userItem.UserName).First(&otherUserItem).Error; err == nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": "User Already Exists"})
			return
		}

		passWord, errHash := components.Hash(userItem.PassWord)

		if errHash != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": "Register Has Failed"})
			return
		}

		userItem.PassWord = passWord

		if err := data.Select("Name", "UserName", "PassWord").Create(&userItem).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": userItem})
	}
}

func GetDetailUserItems(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var userItem models.Users

		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		if err := data.Table(models.Users{}.TableUsers()).
			Where("id = ?", id).
			First(&userItem).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": userItem})
	}
}
