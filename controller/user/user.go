package user

import (
	"net/http"
	"strconv"

	"example.com/m/auth"
	"example.com/m/components"
	"example.com/m/models"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleRegister(data *gorm.DB) gin.HandlerFunc {
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

		var infoUserItem []models.Users

		if err := data.Table(models.Users{}.TableUsers()).
			Where("username = ?", userItem.UserName).First(&infoUserItem).Error; err == nil {
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

func HandleLogin(data *gorm.DB) gin.HandlerFunc {
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

		var infoUserItem models.Users

		if err := data.Table(models.Users{}.TableUsers()).
			Where("username = ?", userItem.UserName).First(&infoUserItem).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": "Username Or Password Incorrect"})
			return
		}

		errHash := components.CheckHash(infoUserItem.PassWord, userItem.PassWord)

		if errHash != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": "Username Or Password Incorrect"})
			return
		}

		token, errCreate := auth.CreateJWT(userItem.UserName)

		if errCreate != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": "Internal Server Error"})
			return
		}

		userDTO := models.UsersDTO{
			Id:        infoUserItem.Id,
			Name:      infoUserItem.Name,
			UserName:  infoUserItem.UserName,
			Status:    infoUserItem.Status,
			Image:     infoUserItem.Image,
			CreatedAt: infoUserItem.CreatedAt,
			UpdatedAt: infoUserItem.UpdatedAt,
			Token:     token,
		}

		userDTO.UsersDTO()

		context.JSON(http.StatusOK, gin.H{"data": userDTO})
	}
}

func GetDetailUserItem(data *gorm.DB) gin.HandlerFunc {
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

func UpdatesUserItem(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		var userItem models.Users

		if err := context.ShouldBind(&userItem); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		if err := data.Where("id = ?", id).
			Updates(&userItem).
			First(&userItem).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": userItem})
	}
}

func DeleteUserItem(data *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		}

		if err := data.Table(models.Users{}.TableUsers()).
			Where("id = ?", id).
			Delete(nil).Error; err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
