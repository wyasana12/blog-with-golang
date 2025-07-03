package handler

import (
	"blog-go/config"
	"blog-go/internal/dto"
	"blog-go/internal/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllRoles(c echo.Context) error {
	var roles []model.Role

	if err := config.DB.Find(&roles).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Failed To Fetch Roles", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, roles)
}

func CreateRole(c echo.Context) error {
	var req dto.CreateRoleRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid Input", "error": err.Error()})
	}

	if err := config.Validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Validation Failed", "error": err.Error()})
	}

	var existing model.Role
	if err := config.DB.Where("name = ?", req.Name).First(&existing).Error; err == nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Name Already In Use"})
	}

	role := model.Role{
		Name: req.Name,
	}

	if err := config.DB.Create(&role).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed To Create Role", "error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Success To Created Role"})
}

func UpdateRole(c echo.Context) error {
	id := c.Param("id")
	roleID, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid Role ID"})
	}

	var req dto.CreateRoleRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid Input"})
	}

	if err := config.Validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Validation Failed", "error": err.Error()})
	}

	var role model.Role

	if err := config.DB.First(&role, roleID).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Role Not Found"})
	}

	role.Name = req.Name
	if err := config.DB.Save(&role).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed To Updated Role", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Success To Updated Role"})
}

func DeleteRole(c echo.Context) error {
	id := c.Param("id")
	roleId, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid Role ID"})
	}

	if err := config.DB.Unscoped().Delete(&model.Role{}, roleId).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed To Deleted Role", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Success To Deleted Role"})
}

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
