package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Product struct {
	ID  bson.ObjectId `bson:"_id"`
	ProductID 			int `json:"product_id"`
	ProductName 		string `json:"product_name"`
	ProductDescription 	string `json:"product_description"`
	ProductIcon 		string `json:"product_icon"`
	Created time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

type ProductCategory struct {
	ID  bson.ObjectId `bson:"_id"`
	ProductCategoryName    		string `json:"product_category_name"`
	ProductCategoryID      		int `json:"product_category_id"`
	ProductCategoryDescription 	string `json:"product_category_description"`
	ProductID					int	`json:"product_id"`
	Created time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

// Products is an array of Products
type Products []Product

// ProductsCategories is an array of BundlesCategories
	type ProductsCategories []ProductCategory
