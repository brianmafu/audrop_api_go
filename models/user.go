package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// User
type User struct {
	ID  bson.ObjectId `bson:"_id"`
	MSISDN string `json:"msisdn"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email string	`json:"email"`
	Created time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

// Users is an array of Users

type Users []User
