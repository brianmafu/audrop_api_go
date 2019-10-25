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
	Songs = "song"
)

func CreateSong(song models.Song) bool {
	// Calling MongoSession from our dbclient.go
	MongoSession.SetMode(mgo.Monotonic, true)
	collection := MongoSession.DB(DBNAME).C(Songs)
	song.ID = bson.NewObjectId()
	song.Created = time.Now()
	song.Modified = time.Now()
	err := collection.Insert(song)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// return list of Songs

func GetSongs() models.Songs {
	// Calling MongoSession from our dbclient.go
	MongoSession.SetMode(mgo.Monotonic, true)
	collection := MongoSession.DB(DBNAME).C(Songs)
	songs := models.Songs{}
	if err := collection.Find(bson.M{}).All(&songs); err != nil {
		log.Println("Failed to read songs:", err)
	}
	log.Println("Success: [GetSongs]")
	return songs
}

func GetSong(songID int) models.Song {
	log.Println("[Getting Song by songID]")
	// Calling MongoSession from our dbclient.go
	MongoSession.SetMode(mgo.Monotonic, true)
	collection := MongoSession.DB(DBNAME).C(Songs)
	song := models.Song{}
	if err := collection.Find(bson.M{"songid": songID}).One(&song); err != nil {
		log.Println("Failed to Read Song:", err)
	}
	log.Println("Success: [GetSong]")
	return song
}

func UpdateSong() {
	log.Println("[Updating Song..]")
}

func DeleteSong() {
	log.Println("[Deleting Song]")
}
