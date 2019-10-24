package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/envy"
	"log"
	"mymtn-shop/docs"
	"mymtn-shop/routes"
)

// @title MyMtn Shop API
// @version 1.0
// @description MyMtn Shop API.

// @host localhost:8085
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
	port := envy.Get("API_PORT", "8085")
	docs.SwaggerInfo.Host = envy.Get("API_HOST", "localhost:"+port)
	log.Fatal(router.Run(fmt.Sprintf(":%s", port)))
}