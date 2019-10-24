package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"mymtn-shop/models"
	"time"
)



const (
	BUNDLES = "bundles"
)



// CreateBundle Inserts Bundle in the DB

func CreateBundle(bundle models.Bundle) bool {
	// Calling MongoSession from our dbclient.go
	MongoSession.SetMode(mgo.Monotonic, true)
	collection := MongoSession.DB(DBNAME).C(BUNDLES)
	bundle.ID = bson.NewObjectId()
	bundle.Created = time.Now()
	bundle.Modified = time.Now()
	err := collection.Insert(bundle)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// GetBundles returns the list of Bundles

func GetBundles() models.Bundles {
	// Calling MongoSession from our dbclient.go
	MongoSession.SetMode(mgo.Monotonic, true)
	collection := MongoSession.DB(DBNAME).C(BUNDLES)
	bundles := models.Bundles{}
	if err := collection.Find(bson.M{}).All(&bundles); err != nil {
		log.Println("Failed to Read Bundles Collection:", err)
	}
	log.Println("Success: [GetBundles]")
	return bundles
}

func GetBundleByBundleID(bundleId int) models.Bundle {
	// Calling MongoSession from our dbclient.go
	MongoSession.SetMode(mgo.Monotonic, true)
	collection := MongoSession.DB(DBNAME).C(BUNDLES)
	bundle := models.Bundle{}
	if err := collection.Find(bson.M{"bundleid": bundleId}).One(&bundle); err != nil {
		log.Println("Failed to Read Bundles Collection:", err)
	}
	log.Println("Success: [GetBundleByBundleID]")
	return bundle
}

func GetBundlesByProductCategoryID(productCategoryId int) models.Bundles {
	log.Println("[Getting Bundle by ProductCategoryID]")
	// Calling MongoSession from our dbclient.go
	MongoSession.SetMode(mgo.Monotonic, true)
	collection := MongoSession.DB(DBNAME).C(BUNDLES)
	bundles := models.Bundles{}
	if err := collection.Find(bson.M{"productcategoryid": productCategoryId}).All(&bundles); err != nil {
		log.Println("Failed to Read Bundles Collection:", err)
	}
	log.Println("Success: [GetBundleByProductCategory]")
	return bundles
}

func GetBundlesByProductID(productId int) models.Bundles {
	log.Println("[Getting Bundle by ProductID]")
	// Calling MongoSession from our dbclient.go
	MongoSession.SetMode(mgo.Monotonic, true)
	collection := MongoSession.DB(DBNAME).C(BUNDLES)
	bundles := models.Bundles{}
	if err := collection.Find(bson.M{"productid": productId}).All(&bundles); err != nil {
		log.Println("Failed to Read Bundles Collection:", err)
	}
	log.Println("Success: [GetBundleByProductID]")
	return bundles
}

func UpdateBundle() {
	log.Println("[Updating Bundle..]")
}

func DeleteBundle()  {
	log.Println("[Deleting Bundle]")
}
