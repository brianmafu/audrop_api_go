package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"mymtn-shop/models"
	"time"
)


// DATABASE COLLECTION NAME the name of the document

const (
	PRODUCTS = "products"
	PRODUCTSCATEGORIES = "productscategories"
)


// CreateProductCategory Insert BundleCategory in the DB

func CreateProductCategory(productCategory models.ProductCategory) bool {
	// Calling MongoSession from our dbclient.go
	MongoSession.SetMode(mgo.Monotonic, true)
	collection := MongoSession.DB(DBNAME).C(PRODUCTSCATEGORIES)
	productCategory.ID = bson.NewObjectId()
	productCategory.Created = time.Now()
	productCategory.Modified = time.Now()
	err := collection.Insert(productCategory)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// GetProductsCategories returns the list of BundleCategories

func GetProductsCategories() models.ProductsCategories {
	// Calling MongoSession from our dbclient.go
	MongoSession.SetMode(mgo.Monotonic, true)
	collection := MongoSession.DB(DBNAME).C(PRODUCTSCATEGORIES)
	productsCategories := models.ProductsCategories{}
	if err := collection.Find(bson.M{}).All(&productsCategories); err != nil {
		log.Println("Failed to Read ProductsCategories:", err)
	}
	log.Println("Success: [GetProductsCategories]")
	return productsCategories
}

func GetProductCategory(productCategoryId int) models.ProductCategory {
	log.Println("[Getting ProductCategory by productCategoryId]")
	// Calling MongoSession from our dbclient.go
	MongoSession.SetMode(mgo.Monotonic, true)
	collection := MongoSession.DB(DBNAME).C(PRODUCTSCATEGORIES)
	productCategory := models.ProductCategory{}
	if err := collection.Find(bson.M{"productcategoryid": productCategoryId}).One(&productCategory); err != nil {
		log.Println("Failed to Read BundlesCategories:", err)
	}
	log.Println("Success: [GetProductCategory]")
	return productCategory
}

// CreateProduct Inserts Product in the DB

func CreateProduct(product models.Product) bool {
	// Calling MongoSession from our dbclient.go
	MongoSession.SetMode(mgo.Monotonic, true)
	collection := MongoSession.DB(DBNAME).C(PRODUCTS)
	product.ID = bson.NewObjectId()
	product.Created = time.Now()
	product.Modified = time.Now()
	err := collection.Insert(product)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// GetBundles returns the list of BundleTypes

func GetProducts() models.Products {
	// Calling MongoSession from our dbclient.go
	MongoSession.SetMode(mgo.Monotonic, true)
	collection := MongoSession.DB(DBNAME).C(PRODUCTS)
	products := models.Products{}
	if err := collection.Find(bson.M{}).All(&products); err != nil {
		log.Println("Failed to read products:", err)
	}
	log.Println("Success: [GetProducts]")
	return products
}

func GetProduct(productId int) models.Product {
	log.Println("[Getting Product by productId]")
	// Calling MongoSession from our dbclient.go
	MongoSession.SetMode(mgo.Monotonic, true)
	collection := MongoSession.DB(DBNAME).C(PRODUCTS)
	product := models.Product{}
	if err := collection.Find(bson.M{"productid": productId}).One(&product); err != nil {
		log.Println("Failed to Read Products:", err)
	}
	log.Println("Success: [GetProduct]")
	return product
}


func UpdateBundleType() {
	log.Println("[Updating BundleType..]")
}

func DeleteBundleType() {
	log.Println("[Deleting BundleType]")
}
