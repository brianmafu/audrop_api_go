package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Song struct {
	ID                bson.ObjectId `bson:"_id"`
	AlbumID           int64         `json:"album_id"`
	ArtistID          int64         `json:"artist_id"`
	Title             string        `json:"title"`
	Description       string        `json:"description"`
	ImageUrl          string        `json:"image_url"`
	ThumbnailImageUrl string        `json:"thumbnail_image_url"`
	Category          string        `json:"category"`
	Duration          string        `json:"duration"`
	DurationInSeconds int64         `json:"duration_in_seconds"`
	Stars             int64         `json:"stars"`
	FileName          string        `json:"file_name"`
}

//   artist: {
//     type: mongoose.Schema.Types.ObjectId,
//     required: true,
//     ref: 'Artist',
//   },
//   title: {
//     type: String,
//     required: true,
//   },
//   description: {
//     type: String,
//     required: true,
//   },
//   imageURL: {
//     type: String,
//     required: true,
//   },
//   thumbnailImageURL: {
//     type: String,
//     required: true,
//   },
//   category: {
//     type: String,
//     required: true,
//   },
//   stars: {
//     type: Number,
//     required: true,
//   },
//   duration: {
//     type: String,
//     required: true,
//   },
//   durationInSeconds: {
//     type: Number,
//     required: true,
//   },
//   fileName: {
//     type: String,
//     required: true,
//   },
// });
