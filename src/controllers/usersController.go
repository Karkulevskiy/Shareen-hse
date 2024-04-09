package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"shareen/src/models"
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

// @Accept json
// @Produce json
// @Success 200 {object} models.Lobby
// @Param userid query string false "Id of user"
// @Router /user/:id [get]
func (uc *UsersController) GetUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	user, responseErr := uc.usersService.GetUser(userId)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// @Accept json
// @Produce json
// @Success 200 {object} models.Lobby
// @Param name query string false "name of user"
// @Router /user/create [post]
func (uc *UsersController) CreateUser(ctx *gin.Context) {
	userName := ctx.Param("name")
	user, responseErr := uc.usersService.CreateUser(&models.User{Name: userName})
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// @Accept json
// @Produce json
// @Success 200 {array} models.Lobby
// @Router /user/allusers [get]
func (uc *UsersController) GetAllUsers(ctx *gin.Context) {
	users, responseErr := uc.usersService.GetAllUsers()
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// @Accept json
// @Produce json
// @Success 204
// @Param userid query string false "Id of user"
// @Router /user/:id [delete]
func (uc *UsersController) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	responseErr := uc.usersService.DeleteUser(userId)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.Status(http.StatusNoContent)
}

// @Accept json
// @Produce json
// @Success 204
// @Param user query string false "Id of user"
// @Router /user/query [patch]
func (uc *UsersController) UpdateUser(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	responseErr := uc.usersService.UpdateUser(&user)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.Status(http.StatusNoContent)
}
