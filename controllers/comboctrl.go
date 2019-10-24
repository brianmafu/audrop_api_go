package controllers

//import "log"
//import "github.com/gin-gonic/gin"
//import "mymtn-shop/models"
//import "strconv"
////import "mymtn-shop/services/dao"
//import "net/http"
//import "encoding/json"
//
//
//func GetComboCtrl(c *gin.Context){
//	ComboId, _ := strconv.Atoi(c.Param("id"))
//	//log.Println(models.GetBundle(ComboId))
//	response := services.GetCombo(int(ComboId))
//	log.Println(response)
//	// Parse Response from apigee and construct JSON using *Bundle struct
//	Combo := models.Combo{}
//	json.Unmarshal([]byte(response), &Combo)
//	c.JSON(http.StatusOK, gin.H{"response": Combo})
//}
//
//func CreateComboCtrl()  {
//	log.Println(models.CreateCombo(1, 10, "1 Week"))
//}
