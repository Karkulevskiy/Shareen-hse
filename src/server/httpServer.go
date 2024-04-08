package server

import (
	"database/sql"
	"log"
	"shareen/src/controllers"
	"shareen/src/repositories"
	"shareen/src/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	router *gin.Engine
	config *viper.Viper
	usersController *controllers.UsersController
	lobbiesController *controllers.LobbiesController
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


	/// etc

	return HttpServer{
		config: config,
		router: router,
		usersController: usersController,
		lobbiesController: lobbiesController,
	}
}

func (hs HttpServer) Start(){
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil{
		log.Fatal("error occured while starting server\n", err)
	}
}