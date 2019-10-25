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

// CreateSong POST
func CreateSongCtrl(c *gin.Context) {
	//  Let gin do the parsing of the body, by using a binding.
	var song models.Song
	_ = c.BindJSON(&song)
	success := dao.CreateSong(song) // Inserts the artist to the DB
	if !success {
		c.JSON(http.StatusCreated, gin.H{"response": "OK"})
		return
	}
}

// GetSongs Controller
func GetSongsCtrl(c *gin.Context) {
	songs := dao.GetAlbums()
	//TODO: Look at refactoring the json.Marshal & Unmarshal [DRY]
	response, err := json.Marshal(songs)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(response, &songs)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "response": songs})
}

// GetSong Controller by songId
func GetSongCtrl(c *gin.Context) {
	songId := c.Param("id")
	id, err := strconv.Atoi(songId)
	if err != nil {
		log.Printf("Error Occur while Converting string %d to %T", id, id)
	}
	song := dao.GetSong(id)
	log.Printf("%+v", song)
	//TODO: Look at refactoring the json.Marshal & Unmarshal [DRY]
	response, err := json.Marshal(song)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(response, &song)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "response": song})
}
