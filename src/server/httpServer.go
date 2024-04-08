package server

import (
	"database/sql"
	"log"
	"shareen/src/controllers"
	_ "shareen/src/docs"
	"shareen/src/repositories"
	"shareen/src/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	Router            *gin.Engine
	config            *viper.Viper
	usersController   *controllers.UsersController
	LobbiesController *controllers.LobbiesController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB, router *gin.Engine) *HttpServer {
	usersRepository := repositories.NewUsersRepository(dbHandler)
	lobbiesRepository := repositories.NewLobbiesRepository(dbHandler)
	usersService := services.NewUsersService(usersRepository, lobbiesRepository)
	lobbiesService := services.NewLobbiesService(usersRepository, lobbiesRepository)
	usersController := controllers.NewUsersController(usersService)
	lobbiesController := controllers.NewLobbiesController(lobbiesService)

	//router.GET("/user/:id", usersController.GetUser) // ??

	/// etc

	return &HttpServer{
		config:            config,
		Router:            router,
		usersController:   usersController,
		LobbiesController: lobbiesController,
	}
}

func (hs *HttpServer) Start() {
	err := hs.Router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatal("error occured while starting server\n", err)
	}
}
