package main

import (
	_ "github.com/swaggo/files"       // swagger embed files
	_ "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"log"
	"shareen/src/configs"
	_ "shareen/src/docs"
	"shareen/src/server"
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
	httpServer := server.InitHttpServer(config, dbHandler)
	httpServer.Start()
}
