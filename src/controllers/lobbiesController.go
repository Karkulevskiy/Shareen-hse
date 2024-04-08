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

func (lc *LobbiesController) CreateLobby(ctx *gin.Context) {
	response, err := lc.lobbiesService.CreateLobby()
	if err != nil {
		log.Println("error occured while creating lobby", err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, response)
}
