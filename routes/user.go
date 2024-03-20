package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/helper"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func register(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data.", "error": err.Error()})
		return
	}
	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user", "error": err.Error()})
		return
	}

	// context.JSON(http.StatusCreated, gin.H{"message": "User created succesfully", "user_detail": user})

	context.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user_detail": gin.H{
			"id":       user.ID,
			"email":    user.Email,
			"password": user.Password, // Catatan: Mengembalikan password hashed
		},
	})

}

func login(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data.", "error": err.Error()})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate user", "error": err.Error()})
		return
	}

	token, err := helper.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not authenticate user", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Login successful with email: %s", user.Email),
		"token":   token,
	})

}
