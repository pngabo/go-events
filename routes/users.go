package routes

import (
	"events-mgmt-portal/models"
	"events-mgmt-portal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUser(ctx *gin.Context) {
	email, err := strconv.ParseInt(ctx.Param("email"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse email."})
		return
	}
	user, err := models.GetEventByID(email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user."})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	user.ID = 1

	err = user.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create user. try again later",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User Created.", "user": user})
}

func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
