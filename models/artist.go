package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Artist struct {
	ID         bson.ObjectId `bson:"_id"`
	ProfileID  int64         `json:"profile_id"`
	StudioName string        `json:"studio_name"`
	FirstName  string        `json:"first_names"`
	LastName   string        `json:"last_name"`
	About      string        `json: "about"`
}

// Artists is an array of Artists
type Artists []Artist

// songs: [
//     {
//       type: mongoose.Schema.Types.ObjectId,
//       ref: 'Song',
//       default: [],
//     },
//   ],
//   albums: [
//     {
//       type: mongoose.Schema.Types.ObjectId,
//       ref: 'Album',
//       default: [],
//     },
//   ],
//   categories: [
//     {
//       type: String,
//       required: true,
//     },
//   ],
//   name: {
//     type: String,
//     required: true,
//   },
//   profileImageURL: {
//     type: String,
//     required: true,
//   },
//   thumbnailProfileImageURL: {
//     type: String,
//     required: true,
//   },
//   about: {
//     type: String,
//     required: true,
//   },
// });
