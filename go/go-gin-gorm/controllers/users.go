package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/magalhaes-andre/database"
	"github.com/magalhaes-andre/models"
)

func GetAllUsers(context *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	context.JSON(http.StatusOK, users)
}

func CreateUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&user)
	context.JSON(http.StatusCreated, user)
}

func UpdateUser(context *gin.Context) {
	var user models.User
	id := context.Param("id")
	database.DB.First(&user, id)

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&user).UpdateColumns(user)
}

func FetchUserById(context *gin.Context) {
	var user models.User
	id := context.Param("id")
	database.DB.First(&user, id)

	if user.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "user not found with id " + id,
		})
		return
	}

	context.JSON(http.StatusOK, user)
}

func FetchByAge(context *gin.Context) {
	var user models.User
	age, err := strconv.Atoi(context.Param("age"))
	if err != nil {
		log.Panic("There was an error when converting age for the following input: ", age, "The error was: ", err)
		return
	}

	database.DB.Where(&models.User{Age: age}).First(&user)

	if user.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "user not found with age " + strconv.Itoa(age),
		})
		return
	}

	context.JSON(http.StatusOK, user)
}

func DeleteUserById(context *gin.Context) {
	var user models.User
	id := context.Param("id")
	database.DB.Delete(&user, id)

	context.JSON(http.StatusOK, gin.H{
		"success": "user with id " + id + " was deleted",
	})
}
