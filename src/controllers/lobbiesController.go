package controllers

import (
	"log"
	"net/http"

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
	response, err := lc.lobbiesService.CreateLobby()
	if err != nil {
		log.Println("error occured while creating lobby", err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, err)
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
	lobby, err := lc.lobbiesService.GetLobby(lobbyId)
	if err != nil {
		log.Println("error getting lobby")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, lobby)
}

// @Accept json
// @Produce json
// @Success 200 {array} models.Lobby
// @Router /lobby/all [get]
func (lc *LobbiesController) GetAllLobbies(ctx *gin.Context) {
	lobbies, err := lc.lobbiesService.GetAllLobbies()
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, lobbies)
}
