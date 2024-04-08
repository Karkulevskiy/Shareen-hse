package main

import (
	"log"
	"shareen/src/configs"
	_ "shareen/src/docs"
	"shareen/src/server"

	//"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

// TODO Посмотреть документацию gin, swagger_gint
func main() {
	log.Println("Starting server")
	log.Print("Initializing configuration")
	config := configs.InitConfig("shareen")
	dbHandler := server.InitDatabase(config)
	router := gin.Default()
	httpServer := server.InitHttpServer(config, dbHandler, router)
	httpServer.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	httpServer.Router.POST("/lobby/create", httpServer.LobbiesController.CreateLobby)
	httpServer.Start()

}
