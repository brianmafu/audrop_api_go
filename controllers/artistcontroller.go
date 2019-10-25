package controllers

import (
	"audrop-api/models"
	"audrop-api/services/dao"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateArtist POST /

func CreateArtistCtrl(c *gin.Context) {
	//  Let gin do the parsing of the body, by using a binding.
	var artist models.Artist
	_ = c.BindJSON(&artist)
	success := dao.CreateArtist(artist) // Inserts the artist to the DB
	if !success {
		c.JSON(http.StatusCreated, gin.H{"response": "OK"})
		return
	}
	//c.JSON(http.StatusOK, c)
}

// GetArtists Controller

func GetArtistsCtrl(c *gin.Context) {
	artists := dao.GetArtists()
	//TODO: Look at refactoring the json.Marshal & Unmarshal [DRY]
	response, err := json.Marshal(artists)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(response, &artists)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "response": artists})
}

// GetArtist Controller by artistid

func GetArtistCtrl(c *gin.Context) {
	artistId := c.Param("id")
	id, err := strconv.Atoi(artistId)
	if err != nil {
		log.Printf("Error Occur while Converting string %d to %T", id, id)
	}
	artist := dao.GetArtist(id)
	log.Printf("%+v", artist)
	//TODO: Look at refactoring the json.Marshal & Unmarshal [DRY]
	response, err := json.Marshal(artist)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(response, &artist)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "response": artist})
}
