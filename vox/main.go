package main

import (
	"log"
	"vox/vox/initializers"
	"vox/vox/services"

	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	server = gin.Default()
}

func main() {

	router := server.Group("/api")

	router.GET("/deneme", services.GetAll)

	log.Fatal(server.Run(":8080"))
}
