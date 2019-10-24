package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Bundle struct {
	ID                bson.ObjectId `bson:"_id"`
	BundleID          int           `json:"bundle_id"`
	ProductID         int           `json:"product_id"`
	BundleValidity    string        `json:"bundle_validity"`
	BundleCost        float64       `json:"bundle_cost"`
	BundleValue       string        `json:"bundle_value"`
	ProductCategoryID int           `json:"product_category_id"`
	Created           time.Time     `json:"created"`
	Modified          time.Time     `json:"modified"`
}

type FilterByCategory struct {
	CatID string `form:"cat_id"`
}

type FilterByProduct struct {
	ProdID string `form:"prod_id"`
}

// Bundles is an array of Bundles
type Bundles []Bundle
