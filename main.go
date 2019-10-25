package main

import (
	"audrop-api/docs"
	"audrop-api/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title Audrop Music API
// @version 1.0
// @description Audrop Music API

// @host localhost:9001
// @BasePath /api/v1/

//Entry point for the Application

func main() {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowWildcard = true
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token")
	router.Use(cors.New(corsConfig))
	routes.ConfigureRouter(router)
	port := os.Getenv("PORT")
	docs.SwaggerInfo.Host = ":" + port
	log.Fatal(router.Run(fmt.Sprintf(":%s", port)))
}
