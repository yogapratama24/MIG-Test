package controllers

import (
	"mitramas_test/helpers"
	"mitramas_test/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service}
}

func (userController *UserController) ReadUserController(c *gin.Context) {
	users, err := userController.service.ReadUser()
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully get users", users).Success(c)
}
