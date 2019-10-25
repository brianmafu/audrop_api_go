package dao

import (
	"audrop-api/models"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DATABASE COLLECTION NAME the name of the document

const (
	Albums = "album"
)

func CreateAlbum(album models.Album) bool {
	// Calling MongoSession from our dbclient.go
	MongoSession.SetMode(mgo.Monotonic, true)
	collection := MongoSession.DB(DBNAME).C(Albums)
	album.ID = bson.NewObjectId()
	album.Created = time.Now()
	album.Modified = time.Now()
	err := collection.Insert(album)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// return list of Albums

func GetAlbums() models.Album {
	// Calling MongoSession from our dbclient.go
	MongoSession.SetMode(mgo.Monotonic, true)
	collection := MongoSession.DB(DBNAME).C(Albums)
	albums := models.Album{}
	if err := collection.Find(bson.M{}).All(&albums); err != nil {
		log.Println("Failed to read albums:", err)
	}
	log.Println("Success: [GetAlbums]")
	return albums
}

func GetAlbum(albumID int) models.Album {
	log.Println("[Getting Album by albumID]")
	// Calling MongoSession from our dbclient.go
	MongoSession.SetMode(mgo.Monotonic, true)
	collection := MongoSession.DB(DBNAME).C(Albums)
	album := models.Album{}
	if err := collection.Find(bson.M{"albumid": albumID}).One(&album); err != nil {
		log.Println("Failed to Read Album:", err)
	}
	log.Println("Success: [GetAlbum]")
	return album
}

func UpdateAlbum() {
	log.Println("[Updating Album..]")
}

func DeleteAlbum() {
	log.Println("[Deleting Album]")
}
