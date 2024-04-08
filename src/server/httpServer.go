package server

import (
	"database/sql"
	"shareen/src/controllers"
	"shareen/src/repositories"
	"shareen/src/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	router *gin.Engine
	config *viper.Viper
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer{
	usersRepository := repositories.NewUsersRepository(dbHandler)
	lobbiesRepository := repositories.NewLobbiesRepository(dbHandler)
	usersService := services.NewUsersService(usersRepository, lobbiesRepository)
	lobbiesService := services.NewLobbiesService(usersRepository, lobbiesRepository)
	usersController := controllers.NewUsersController(usersService)
	lobbiesController := controllers.NewLobbiesController(lobbiesService)
	router := gin.Default()
	router.POST("/lobby", lobbiesController.CreateLobby)
	router.GET("/user/:id", usersController.GetUser) // ??
}
