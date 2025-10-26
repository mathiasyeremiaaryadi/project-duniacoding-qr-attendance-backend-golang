package controllers

import (
	"net/http"
	"qr-attendance-backend/databases"
	"qr-attendance-backend/helpers"
	"qr-attendance-backend/models"
	"qr-attendance-backend/structs"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	req := structs.UserLoginRequest{}
	user := models.User{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorsResponse{
			Success: false,
			Message: "Validasi error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := databases.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorsResponse{
			Success: false,
			Message: "invalid login",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorsResponse{
			Success: false,
			Message: "invalid login",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	token, err := helpers.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorsResponse{
			Success: false,
			Message: "failed create token",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "Login Success",
		Data: structs.UserResponse{
			Id:        uint(user.ID),
			Name:      user.Name,
			Username:  user.Username,
			RoleId:    user.RoleId,
			NoReg:     user.NoReg,
			Email:     user.Email,
			Token:     &token,
			CreatedAt: user.CreatedAt.Format("15:04:05 02-01-2006"),
			UpdatedAt: user.UpdatedAt.Format("15:04:05 02-01-2006"),
		},
	})
}
