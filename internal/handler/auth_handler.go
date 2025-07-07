package handler

import (
	"blog-go/config"
	"blog-go/helper"
	"blog-go/internal/dto"
	"blog-go/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	var req dto.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid Input"})
	}

	if req.Password != req.PasswordConfirm {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Password Not Matched"})
	}

	v := helper.CustomValidation()
	if err := v.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Validation Failed", "error": err.Error()})
	}

	var existing model.User
	if err := config.DB.Where("email = ?", req.Email).First(&existing).Error; err == nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Email Already In Use"})
	}

	hash, _ := helper.HashPassword(req.Password)

	user := model.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: hash,
	}

	var role model.Role
	if err := config.DB.Where("name = ?", "author").FirstOrCreate(&role).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Default Role Not Found"})
	}

	user.Roles = []model.Role{role}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Registration Failed"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Registration Success"})
}

func Login(c echo.Context) error {
	var req dto.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid Input"})
	}

	var user model.User
	if err := config.DB.Preload("Roles").Where("email = ? OR username = ?", req.Email, req.Username).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Email Is Not Valid"})
	}

	if err := helper.VerifyPassword(user.Password, req.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Password Has Not Matched"})
	}

	var roles []string
	for _, role := range user.Roles {
		roles = append(roles, role.Name)
	}

	token, _ := helper.CreateToken(user.ID, roles)

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
		"user": echo.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"roles": roles,
		}})
}

func Me(c echo.Context) error {
	user := c.Get("user").(model.User)
	return c.JSON(http.StatusOK, user)
}
