package routes

import (
	"audrop-api/controllers"
	_ "audrop-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRouter(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("health", controllers.HealthCheck)

		apiV1.POST("albums", controllers.CreateAlbumCtrl)
		apiV1.POST("songs", controllers.CreateSongCtrl)
		apiV1.POST("artists", controllers.CreateArtistCtrl)

		apiV1.POST("users", controllers.CreateUserCtrl)
		apiV1.GET("users", controllers.GetUsersCtrl)
		apiV1.GET("users/:msisdn", controllers.GetUserCtrl)

		apiV1.GET("artists", controllers.GetArtistsCtrl)
		apiV1.GET("artists/:id", controllers.GetArtistCtrl)

		apiV1.GET("albums", controllers.GetAlbumsCtrl)
		apiV1.GET("albums/:id", controllers.GetAlbumCtrl)

		apiV1.GET("songs", controllers.GetSongsCtrl)
		apiV1.GET("songs/:id", controllers.GetSongCtrl)

	}
}
