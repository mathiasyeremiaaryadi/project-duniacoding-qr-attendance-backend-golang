package controllers

import (
	"net/http"
	"qr-attendance-backend/databases"
	"qr-attendance-backend/helpers"
	"qr-attendance-backend/models"
	"qr-attendance-backend/structs"

	"github.com/gin-gonic/gin"
)

func ListUser(c *gin.Context) {
	var users []models.User

	databases.DB.Find(&users)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "List data users",
		Data:    users,
	})
}

func CreateUser(c *gin.Context) {
	req := structs.UserCreateRequest{}

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
		Email:    req.Email,
		RoleId:   req.RoleId,
		NoReg:    req.NoReg,
		Password: helpers.HashPassword(req.Password),
	}

	if err := databases.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorsResponse{
			Success: false,
			Message: "failed to create user",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "User created successfully",
		Data: structs.UserResponse{
			Name:     user.Name,
			Username: user.Email,
			RoleId:   user.RoleId,
			NoReg:    user.NoReg,
		},
	})
}

func UserById(c *gin.Context) {
	var user models.User

	userId := c.Param("id")

	databases.DB.Where("id = ?", userId).First(&user)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "List data users",
		Data:    user,
	})
}

func UpdateUser(c *gin.Context) {
	userId := c.Param("id")
	user := models.User{}

	req := structs.UserUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorsResponse{
			Success: false,
			Message: "Validasi error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := databases.DB.First(&user, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorsResponse{
			Success: false,
			Message: "User not found",
		})
		return
	}

	user.Name = req.Name
	user.Username = req.Username
	user.Email = req.Email
	user.RoleId = req.RoleId
	user.Password = req.Password

	if err := databases.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorsResponse{
			Success: false,
			Message: "failed to create user",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "User created successfully",
		Data: structs.UserResponse{
			Name:     user.Name,
			Username: user.Email,
			RoleId:   user.RoleId,
			NoReg:    user.NoReg,
		},
	})
}

func DeleteUser(c *gin.Context) {
	var user models.User

	userId := c.Param("id")

	if err := databases.DB.First(&user, userId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorsResponse{
			Success: false,
			Message: "user not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := databases.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorsResponse{
			Success: false,
			Message: "failed to delete user",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "List data users",
		Data:    user,
	})
}
