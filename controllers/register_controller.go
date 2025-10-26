package controllers

import (
	"net/http"
	"qr-attendance-backend/databases"
	"qr-attendance-backend/helpers"
	"qr-attendance-backend/models"
	"qr-attendance-backend/structs"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req = structs.UserCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorsResponse{
			Success: false,
			Message: "Validasi error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	user := models.User{
		Name:     req.Name,
		Username: req.Username,
		NoReg:    req.NoReg,
		Email:    req.Email,
		RoleId:   req.RoleId,
		Password: helpers.HashPassword(req.Password),
	}

	if err := databases.DB.Create(&user).Error; err != nil {
		if helpers.IsDupliateEntryError(err) {
			c.JSON(http.StatusConflict, structs.ErrorsResponse{
				Success: false,
				Message: "Duplicate entry error",
				Errors:  helpers.TranslateErrorMessage(err),
			})
		} else {
			c.JSON(http.StatusInternalServerError, structs.ErrorsResponse{
				Success: false,
				Message: "failed to register user",
				Errors:  helpers.TranslateErrorMessage(err),
			})
		}
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "User created successfully",
		Data: structs.UserResponse{
			Id:        uint(user.ID),
			Name:      user.Name,
			Username:  user.Username,
			RoleId:    user.RoleId,
			NoReg:     user.NoReg,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format("15:04:05 02-01-2006"),
			UpdatedAt: user.UpdatedAt.Format("15:04:05 02-01-2006"),
		},
	})
}
