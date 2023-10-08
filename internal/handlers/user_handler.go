package handlers

import (
	"net/http"

	"github.com/frangar97/testapi/internal/models"
	"github.com/labstack/echo/v4"
)

func (h Handlers) CreateUserHandler(c echo.Context) error {
	userModel := models.UserModel{}
	err := c.Bind(&userModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, requestResponse{Message: "user and password are required"})
	}

	err = h.services.UserService.CreateUser(userModel.User, userModel.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, requestResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, requestResponse{Message: "user created"})
}

func (h Handlers) LoginUserHandler(c echo.Context) error {
	userModel := models.UserModel{}
	err := c.Bind(&userModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, requestResponse{Message: err.Error()})
	}

	token, err := h.services.UserService.LoginUser(userModel.User, userModel.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, requestResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, requestResponse{Message: token})
}
