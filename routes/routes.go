package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"mymtn-shop/controllers"
	_ "mymtn-shop/docs"
)

func ConfigureRouter(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("health", controllers.HealthCheck)

		apiV1.POST("products", controllers.CreateProductCtrl)
		apiV1.POST("categories", controllers.CreateProductCategoryCtrl)
		apiV1.POST("bundles", controllers.CreateBundleCtrl)

		apiV1.POST("users", controllers.CreateUserCtrl)
		apiV1.GET("users", controllers.GetUsersCtrl)
		apiV1.GET("users/:msisdn", controllers.GetUserCtrl)

		apiV1.GET("products", controllers.GetProductsCtl)
		apiV1.GET("products/:id", controllers.GetProductCtrl)
		apiV1.GET("categories", controllers.GetProductsCategoriesCtl)
		apiV1.GET("categories/:id", controllers.GetProductCategoryCtrl)
		apiV1.GET("bundles", controllers.GetBundlesCtrl)
		apiV1.GET("bundles/:id", controllers.GetBundleCtrl)

		apiV1.GET("product-bundles/:prod_id", controllers.GeProductBundlesCtrl)
		apiV1.GET("product-cat-bundles/:cat_id", controllers.GeProductCategoryBundlesCtrl)

		//apiV1.GET("combos/:id", controllers.GetComboCtrl)
	}
}