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
