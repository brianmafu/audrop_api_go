package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"mymtn-shop/models"
	"time"
)

// Repository ...
//type Repository struct{}

// DATABASE COLLECTION NAME the name of the document

const USERS = "users"

// GetUsers returns the list of Users

func GetUsers() models.Users {
	// Calling MongoSession from our dbclient.go
	MongoSession.SetMode(mgo.Monotonic, true)
	collection := MongoSession.DB(DBNAME).C(USERS)
	users := models.Users{}
	if err := collection.Find(bson.M{}).All(&users); err != nil {
		log.Println("Failed to write users:", err)
	}
	log.Println("Success: [GetUsers]")
	return users
}

func GetUser(msisdn string) models.User {
	log.Println("[Getting User by MSISDN]")
	// Calling MongoSession from our dbclient.go
	MongoSession.SetMode(mgo.Monotonic, true)
	collection := MongoSession.DB(DBNAME).C(USERS)
	user := models.User{}
	if err := collection.Find(bson.M{"msisdn": msisdn}).One(&user); err != nil {
		log.Println("Failed to Read users:", err)
	}
	log.Println("Success: [GetUser]")
	return user
}

// CreateUser Inserts User in the DB

func CreateUser(user models.User) bool {
	// Calling MongoSession from our dbclient.go
	MongoSession.SetMode(mgo.Monotonic, true)
	collection := MongoSession.DB(DBNAME).C(USERS)
	user.ID = bson.NewObjectId()
	user.Created = time.Now()
	user.Modified = time.Now()
	err := collection.Insert(user)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func UpdateUser() {
	log.Println("[Updating User..]")
}

func DeleteUser()  {
	log.Println("[Deleting User]")
}
