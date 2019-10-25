package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"audrop-api/models"
	"audrop-api/services/dao"
	"net/http"
)

// CreateUser POST /

func CreateUserCtrl(c *gin.Context) {

	// This is for your learning process only
	//x, _ := ioutil.ReadAll(c.Request.Body)
	//log.Printf("%s", string(x))
	//c.JSON(http.StatusOK, c)

	//  Let gin do the parsing of the body, by using a binding.
	var user models.User
	_ = c.BindJSON(&user)
	success := dao.CreateUser(user) // Inserts the user to the DB
	log.Printf("%+v", user)
	if !success {
		c.JSON(http.StatusCreated, gin.H{"response": "OK"})
		return
	}
	//c.JSON(http.StatusOK, c)
}

func GetUsersCtrl(c *gin.Context)  {
	users := dao.GetUsers()
	//TODO: Look at refactoring the json.Marshal & Unmarshal [DRY]
	response, err := json.Marshal(users)
	if err != nil {panic(err)}
	err = json.Unmarshal(response, &users)
	if err != nil {panic(err)}
	c.JSON(http.StatusOK, gin.H{"statusCode" : http.StatusOK, "response": users})
}

func GetUserCtrl(c *gin.Context){
	msisdn := c.Param("msisdn")
	user := dao.GetUser(msisdn)
	response, err := json.Marshal(user)
	if err != nil {panic(err)}
	err = json.Unmarshal(response, &user)
	if err != nil {panic(err)}
	c.JSON(http.StatusOK, gin.H{"statusCode" : http.StatusOK, "response": user})
}