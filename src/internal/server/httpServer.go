package server

import (
	"database/sql"
	"log"
	controllers "shareen/src/internal/api"
	_ "shareen/src/internal/docs"
	"shareen/src/internal/repositories"
	"shareen/src/internal/services"

	"github.com/gin-contrib/cors"
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

	router.Use(cors.Default())                                                //Allow all origing | Нужно будет потом настроить
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler)) // Using swagger

	lobbyGroup := router.Group("/lobby")
	{
		lobbyGroup.GET(":id", lobbiesController.GetLobby)
		lobbyGroup.GET("all", lobbiesController.GetAllLobbies)
		lobbyGroup.GET("lobbyusers/:id", lobbiesController.GetLobbyUsers)
		lobbyGroup.POST("create", lobbiesController.CreateLobby)
		lobbyGroup.POST("seturl", lobbiesController.SetVideoURL)
		lobbyGroup.DELETE("delete/:id", lobbiesController.DeleteLobby)
		lobbyGroup.PATCH("update", lobbiesController.UpdateLobby)
	}

	userGroup := router.Group("/user")
	{
		userGroup.GET(":id", usersController.GetUser)
		userGroup.GET("allusers", usersController.GetAllUsers)
		userGroup.POST("create/:name", usersController.CreateUser)
		userGroup.POST("join", usersController.JoinUserInLobby)
		userGroup.PATCH("update", usersController.UpdateUser)
		userGroup.DELETE(":id", usersController.DeleteUser)
	}

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
