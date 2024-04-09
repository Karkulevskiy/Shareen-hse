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
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type HttpServer struct {
	router            *gin.Engine
	config            *viper.Viper
	usersController   *controllers.UsersController
	lobbiesController *controllers.LobbiesController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	usersRepository := repositories.NewUsersRepository(dbHandler)
	lobbiesRepository := repositories.NewLobbiesRepository(dbHandler)
	usersService := services.NewUsersService(usersRepository, lobbiesRepository)
	lobbiesService := services.NewLobbiesService(usersRepository, lobbiesRepository)
	usersController := controllers.NewUsersController(usersService)
	lobbiesController := controllers.NewLobbiesController(lobbiesService)

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	lobbyGroup := router.Group("/lobby")
	{
		lobbyGroup.GET("/:id", lobbiesController.GetLobby)
		lobbyGroup.POST("/create", lobbiesController.CreateLobby)
		lobbyGroup.GET("/all", lobbiesController.GetAllLobbies)
	}

	//router.GET("/user/:id", usersController.GetUser) // ??

	/// etc

	return HttpServer{
		config:            config,
		router:            router,
		usersController:   usersController,
		lobbiesController: lobbiesController,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatal("error occured while starting server\n", err)
	}
}
