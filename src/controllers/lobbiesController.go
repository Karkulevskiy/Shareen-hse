package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"shareen/src/models"
	"shareen/src/services"

	"github.com/gin-gonic/gin"
)

type LobbiesController struct {
	lobbiesService *services.LobbiesService
}

func NewLobbiesController(lobbiesService *services.LobbiesService) *LobbiesController {
	return &LobbiesController{
		lobbiesService: lobbiesService,
	}
}

// @Accept json
// @Produce json
// @Success 200 {object} models.Lobby
// @Router /lobby/create [post]
func (lc *LobbiesController) CreateLobby(ctx *gin.Context) {
	response, responseErr := lc.lobbiesService.CreateLobby()
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// @Accept json
// @Produce json
// @Success 200 {object} models.Lobby
// @Param lobbyid query string false "Id of lobby"
// @Router /lobby/:id [get]
func (lc *LobbiesController) GetLobby(ctx *gin.Context) {
	lobbyId := ctx.Param("id")
	lobby, responseErr := lc.lobbiesService.GetLobby(lobbyId)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, lobby)
}

// @Accept json
// @Produce json
// @Success 200 {array} models.Lobby
// @Router /lobby/all [get]
func (lc *LobbiesController) GetAllLobbies(ctx *gin.Context) {
	lobbies, responseErr := lc.lobbiesService.GetAllLobbies()
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, lobbies)
}

// @Accept json
// @Success 204
// @Failure 500
// @Router /lobby/delete [delete]
func (lc *LobbiesController) DeleteLobby(ctx *gin.Context) {
	lobbyId := ctx.Param("id")
	responseErr := lc.lobbiesService.DeleteLobby(lobbyId)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.Status(http.StatusNoContent)
}

// @Accept json
// @Success 204
// @Failure 500
// @Router /lobby/deleteall [delete]
func (lc *LobbiesController) DeleteAllLobbies(ctx *gin.Context) {
	responseErr := lc.lobbiesService.DeleteAllLobbies()
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.Status(http.StatusNoContent)
}

// @Accept json
// @Success 204
// @Failure 500
// @Router /lobby/update [patch]
// @Param lobby query string false
func (lc *LobbiesController) UpdateLobby(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var lobby models.Lobby
	err = json.Unmarshal(body, &lobby)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	responseErr := lc.lobbiesService.UpdateLobby(&lobby)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.Status(http.StatusNoContent)
}

// @Accept json
// @Success 200
// @Failure 500
// @Router /lobby/lobbyusers [PATCH]
// @Param lobby query string false
func (lc *LobbiesController) GetLobbyUsers(ctx *gin.Context) {
	lobbyId := ctx.Param("id")
	users, responseErr := lc.lobbiesService.GetLobbyUsers(lobbyId)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, users)
}
