package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nullifidianz/go-crud/src/controller/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server := gin.Default()

	routes.InitRoutes(&server.RouterGroup)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
