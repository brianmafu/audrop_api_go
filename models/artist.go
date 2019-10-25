package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Artist struct {
	ID         bson.ObjectId `bson:"_id"`
	ProfileID  int64         `json:"profile_id"`
	StudioName string        `json:"studio_name"`
	FirstName  string        `json:"first_names"`
	LastName   string        `json:"last_name"`
	About      string        `json: "about"`
	Created    time.Time     `json:"created"`
	Modified   time.Time     `json:"modified"`
}

// Array of Artists
type Artists []Artist
