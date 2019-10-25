package dao

import (
	"audrop-api/models"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DATABASE COLLECTION NAME the name of the document

const (
	Artists = "artists"
)

func CreateArtist(artist models.Artist) bool {
	// Calling MongoSession from our dbclient.go
	fmt.Println("on CreateArtist")
	MongoSession.SetMode(mgo.Monotonic, false)

	// if error != nil {
	// 	fmt.Println("some mognn conntect bug")
	// 	fmt.Println(error)
	// 	fmt.Println("*********")
	// }
	fmt.Println(DBNAME)
	fmt.Println(artist)
	collection := MongoSession.DB(DBNAME).C(Artists)
	artist.ID = bson.NewObjectId()
	artist.Created = time.Now()
	artist.Modified = time.Now()
	err := collection.Insert(artist)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// return list of Artists

func GetArtists() models.Artist {
	// Calling MongoSession from our dbclient.go
	MongoSession.SetMode(mgo.Monotonic, true)
	collection := MongoSession.DB(DBNAME).C(Artists)
	artists := models.Artist{}
	if err := collection.Find(bson.M{}).All(&artists); err != nil {
		log.Println("Failed to read artists:", err)
	}
	log.Println("Success: [GetArtists]")
	return artists
}

func GetArtist(artistID int) models.Artist {
	log.Println("[Getting Artist by artistID]")
	// Calling MongoSession from our dbclient.go
	MongoSession.SetMode(mgo.Monotonic, true)
	collection := MongoSession.DB(DBNAME).C(Artists)
	artist := models.Artist{}
	if err := collection.Find(bson.M{"artistid": artistID}).One(&artist); err != nil {
		log.Println("Failed to Read Artist:", err)
	}
	log.Println("Success: [GetArtist]")
	return artist
}

func UpdateArtist() {
	log.Println("[Updating Artist..]")
}

func DeleteArtist() {
	log.Println("[Deleting Artist]")
}
