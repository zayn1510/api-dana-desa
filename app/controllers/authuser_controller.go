package controllers

import (
	"apidanadesa/app/requests"
	"apidanadesa/app/resources"
	"apidanadesa/app/services"
	"github.com/gin-gonic/gin"
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
		resources.BadRequest(c, err)
		return
	}
	if err := uc.service.CreateUser(r); err != nil {
		resources.InternalError(c, err)
		return
	}
	resources.Success(c, "user created")
}
func (uc *AuthUserController) Login(c *gin.Context) {
	var r requests.UserRequestLogin
	if err := c.ShouldBindJSON(&r); err != nil {
		resources.BadRequest(c, err)
		return
	}
	token, err := uc.service.LoginUser(r)
	if err != nil {
		resources.InternalError(c, err)
		return
	}
	resources.Success(c, "user logged in", token)
}
