package handler

import (
	"blog-go/config"
	"blog-go/internal/dto"
	"blog-go/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AssignRoleToUser(c echo.Context) error {
	userID := c.Param("id")
	var req dto.AssignRoleRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid Input", "error": err.Error()})
	}

	if err := config.Validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Validation Failed", "error": err.Error()})
	}

	var user model.User
	if err := config.DB.Preload("Roles").First(&user, userID).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "User Not Found"})
	}

	var role model.Role
	if err := config.DB.Where("name = ?", req.RoleName).First(&role).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Role Not Found"})
	}

	for _, r := range user.Roles {
		if r.ID == role.ID {
			return c.JSON(http.StatusConflict, echo.Map{"message": "User Already Has Role"})
		}
	}

	if err := config.DB.Model(&user).Association("Roles").Append(&role); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed To Assign Role", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Role Assigned Succesfully"})
}

func RevokeRoleFromUser(c echo.Context) error {
	userID := c.Param("id")
	roleID := c.Param("roleId")

	var user model.User
	if err := config.DB.Preload("Roles").First(&user, userID).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "User Not Found"})
	}

	var role model.Role
	if err := config.DB.First(&role, roleID).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Role Not Found"})
	}

	if err := config.DB.Model(&user).Association("Roles").Delete(&role); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed To Revoke Role", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Role Revoked Succesfully"})
}
