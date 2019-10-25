package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Album struct {
	ID                bson.ObjectId `bson:"_id"`
	ArtistID          int64         `json:"artist_id"`
	Title             string        `json:"title"`
	Description       string        `json:"description"`
	ImageUrl          string        `json:"image_url"`
	ThumbnailImageUrl string        `json:"thumbnail_image_url"`
	Category          string        `json:"category"`
	Duration          string        `json:"duration"`
	DurationInSeconds int64         `json:"duration_in_seconds"`
	Stars             int64         `json:"stars"`
	Created           time.Time     `json:"created"`
	Modified          time.Time     `json:"modified"`
}

// Array of Albums
type Albums []Album
