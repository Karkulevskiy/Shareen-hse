package controllers

import (
	"net/http"
	"shareen/src/services"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	usersService *services.UsersService
}

func NewUsersController(usersService *services.UsersService) *UsersController {
	return &UsersController{
		usersService: usersService,
	}
}

func (uc *UsersController) GetUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	response, err := uc.usersService.GetUser(userId)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, response)
}
