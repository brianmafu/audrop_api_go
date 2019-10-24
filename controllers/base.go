package controllers

// import (
// 	"fmt"
// 	"mymtn-shop/models"
// 	"mymtn-shop/services"
// 	"mymtn-shop/testData/services"
// 	"net/http"
// 	"encoding/json"
// 	"strconv"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// // HealthCheck godoc
// //// @Summary Healthcheck endpoint for Shop API
// //// @Description This endpoint exists solely for checking the active status of the application
// //// @Description Any Http status other than 200 signifies that the application is down
// //// @Accept json
// //// @Produce json
// //// @Success 200
// //// @Router /health [get]
// func HealthCheck(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{"message": "Status is ok."})
// }

// // TestRequest godoc
// // @Summary Test Request endpoint to get data from some backend service
// // @Description This endpoint exists testing sample request calls to backing service
// // @Accept json
// // @Produce json
// // @Success 200
// // @Router /testdata/{number} [get]
// func TestRequest(c *gin.Context) {
// 	number, _ := strconv.Atoi(c.Param("number"))
// 	result, err := services.GetSampleData(number)
// 	if err != nil {
// 		c.JSON(http.StatusExpectationFailed, gin.H{"error": "Something went wrong."})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"result": result})
// }
// func CreateUser(c *gin.Context) {
// 	var user models.User

// 	_ = c.BindJSON(&user)
// 	go func(user models.User) {
// 		time.Sleep(time.Second * 10) //Simulating long running process...
// 		fmt.Println("Doing something here")
// 	}(user)
// 	c.JSON(http.StatusOK, gin.H{"message": "Your request is being processed"})

// }

// func GetProductTypes(c *gin.Context) {
// 	product_types, error := testData.GetProductTypes()
// 	if error != nil {
// 		c.JSON(http.StatusExpectationFailed, gin.H{"error": "Something went wrong"})
// 		return
// 	}
// 	var response map[string]interface{}

//     json.Unmarshal(product_types, &response)
// 	c.JSON(http.StatusOK, gin.H{"message": response})
// }

// func GetProducts(c *gin.Context) {
// 	products, error := testData.GetProducts()
// 	if error != nil {
// 		c.JSON(http.StatusExpectationFailed, gin.H{"error": "Something went wrong"})
// 		return
// 	}
// 	var prods models.Products

// 	err := json.Unmarshal(products, &prods)
// 	if err != nil{
// 		c.JSON(http.StatusBadRequest, gin.H{"error":err})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": prods})
// }
