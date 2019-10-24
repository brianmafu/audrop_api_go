package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"mymtn-shop/models"
	"mymtn-shop/services/dao"
	"net/http"
	"strconv"
)

// CreateBundle POST /

func CreateBundleCtrl(c *gin.Context) {

	//  Let gin do the parsing of the body, by using a binding.
	var bundle models.Bundle
	_ = c.BindJSON(&bundle)
	success := dao.CreateBundle(bundle) // Inserts the bundle to the DB
	log.Printf("%+v", bundle)
	if !success {
		c.JSON(http.StatusCreated, gin.H{"response": "OK"})
		return
	}
	//c.JSON(http.StatusOK, c)
}

func GetBundleCtrl(c *gin.Context){
	bundleId := c.Param("id")
	id, err := strconv.Atoi(bundleId)
	if err != nil {log.Printf("Error Occur while Converting string %d to %T", id, id)}
	bundle := dao.GetBundleByBundleID(id)
	//TODO: Look at refactoring the json.Marshal & Unmarshal [DRY]
	response, err := json.Marshal(bundle)
	if err != nil {panic(err)}
	err = json.Unmarshal(response, &bundle)
	if err != nil {panic(err)}
	c.JSON(http.StatusOK, gin.H{"statusCode" : http.StatusOK, "response": bundle})
}

func GetBundlesCtrl(c *gin.Context)  {
	bundles := dao.GetBundles()
	//TODO: Look at refactoring the json.Marshal & Unmarshal [DRY]
	response, err := json.Marshal(bundles)
	if err != nil {panic(err)}
	err = json.Unmarshal(response, &bundles)
	if err != nil {panic(err)}
	if len(bundles) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "response": "No Bundles found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"statusCode" : http.StatusOK, "response": bundles})
	return
}

func GeProductBundlesCtrl(c *gin.Context) {
	prodId := c.Param("prod_id")
	id, err := strconv.Atoi(prodId)
	bundles := dao.GetBundlesByProductID(id)
	response, err := json.Marshal(bundles)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(response, &bundles)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "response": bundles})
	return

}

func GeProductCategoryBundlesCtrl(c *gin.Context) {
	catId := c.Param("cat_id")
	id, err := strconv.Atoi(catId)
	bundles := dao.GetBundlesByProductCategoryID(id)
	response, err := json.Marshal(bundles)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(response, &bundles)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "response": bundles})
	return
}

