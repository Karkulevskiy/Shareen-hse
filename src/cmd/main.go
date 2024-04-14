package main

import (
	_ "github.com/swaggo/files"       // swagger embed files
	_ "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"log"
	"shareen/src/internal/configs"
	_ "shareen/src/internal/docs"
	"shareen/src/internal/server"
)

// @title           Swagger Shareen
// @version         1.0

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI

// TODO Посмотреть документацию gin, swagger_gint
func main() {
	log.Println("Starting server")
	log.Print("Initializing configuration")
	config := configs.InitConfig("shareen")
	dbHandler := server.InitDatabase(config)
	httpServer := server.InitHttpServer(config, dbHandler)
	httpServer.Start()
}
