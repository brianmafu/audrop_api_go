package dao

import (
	"time"

	"gopkg.in/mgo.v2"
)

var (
	// MongoDBHosts = os.Getenv("DB_HOST")
	// DBNAME = os.Getenv("DB_NAME")
	// DBPASSWORD = os.Getenv("DB_PASSWORD")
	// DBUSERNAME = ""

	MongoDBHosts = "cluster0-wwqx1.mongodb.net"
	DBNAME       = "audrop"
	DBPASSWORD   = "audrop"
	DBUSERNAME   = "audrop"

	// We need this object to establish a session to our MongoDB.
	mongoDBDialInfo = &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: DBNAME,
		Username: DBUSERNAME,
		Password: DBPASSWORD,
	}

	// Create a session which maintains a pool of socket connections to our MongoDB.
	MongoSession, err = mgo.DialWithInfo(mongoDBDialInfo)
)

//func int() {
//	if err != nil {
//		log.Fatalf("Failed to establish Connection to Mongo server: %s", err)
//	}
//	log.Println("Successfully Connect to Mongo server:", MongoDBHosts)
//}
