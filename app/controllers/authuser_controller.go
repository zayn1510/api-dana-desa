package controllers

import (
	"apidanadesa/app/requests"
	"apidanadesa/app/resources"
	"apidanadesa/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthUserController struct {
	service *services.UsersService
}

func NewAuthUserController() *AuthUserController {
	return &AuthUserController{
		service: services.NewUsersService(),
	}
}

func (uc *AuthUserController) RegisterUser(c *gin.Context) {
	var r requests.UserRequestCreate
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}
	if err := uc.service.CreateUser(r); err != nil {
		c.JSON(http.StatusInternalServerError, resources.Response{
			Message: "internal server error",
			Status:  false,
		})
		return
	}
	c.JSON(http.StatusOK, resources.Response{
		Message: "create user successfully",
		Status:  true,
	})
}
func (uc *AuthUserController) Login(c *gin.Context) {
	var r requests.UserRequestLogin
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}
	token, err := uc.service.LoginUser(r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resources.Response{
			Message: err.Error(),
			Status:  false,
		})
		return
	}
	c.JSON(http.StatusOK, resources.Response{
		Message: "login successfully",
		Status:  true,
		Data:    token,
	})
}
