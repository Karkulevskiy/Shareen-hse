package main

import (
	"log"
	"shareen/src/configs"
	_ "shareen/src/docs"
	"shareen/src/server"
)

//TODO Посмотреть документацию gin, swagger_gin

func main() {
	log.Println("Starting server")
	log.Print("Initializing configuration")
	config := configs.InitConfig("shareen")
	dbHandler := server.InitDatabase(config)
	httpServer := server.InitHttpServer(config, dbHandler)
	httpServer.Start()
}
