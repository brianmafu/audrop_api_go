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

// CreateAlbum POST /

func CreateAlbumCtrl(c *gin.Context) {
	//  Let gin do the parsing of the body, by using a binding.
	var album models.Album
	_ = c.BindJSON(&album)
	success := dao.CreateAlbum(album) // Inserts the artist to the DB
	if !success {
		c.JSON(http.StatusCreated, gin.H{"response": "OK"})
		return
	}
	//c.JSON(http.StatusOK, c)
}

// GetAlbums Controller

func GetAlbumsCtrl(c *gin.Context) {
	albums := dao.GetAlbums()
	//TODO: Look at refactoring the json.Marshal & Unmarshal [DRY]
	response, err := json.Marshal(albums)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(response, &albums)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "response": albums})
}

// GetAlbum Controller by artistid

func GetAlbumCtrl(c *gin.Context) {
	albumId := c.Param("id")
	id, err := strconv.Atoi(albumId)
	if err != nil {
		log.Printf("Error Occur while Converting string %d to %T", id, id)
	}
	album := dao.GetAlbum(id)
	log.Printf("%+v", album)
	//TODO: Look at refactoring the json.Marshal & Unmarshal [DRY]
	response, err := json.Marshal(album)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(response, &album)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "response": album})
}
