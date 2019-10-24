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


// CreateProduct POST /

func CreateProductCtrl(c *gin.Context) {
	//  Let gin do the parsing of the body, by using a binding.
	var product models.Product
	_ = c.BindJSON(&product)
	success := dao.CreateProduct(product) // Inserts the product to the DB
	if !success {
		c.JSON(http.StatusCreated, gin.H{"response": "OK"})
		return
	}
	//c.JSON(http.StatusOK, c)
}

// Home Screen Method

func GetProductsCtl(c *gin.Context)  {
	products := dao.GetProducts()
	//TODO: Look at refactoring the json.Marshal & Unmarshal [DRY]
	response, err := json.Marshal(products)
	if err != nil {panic(err)}
	err = json.Unmarshal(response, &products)
	if err != nil {panic(err)}
	c.JSON(http.StatusOK, gin.H{"statusCode" : http.StatusOK, "response": products})
}

func GetProductCtrl(c *gin.Context){
	productId := c.Param("id")
	id, err := strconv.Atoi(productId)
	if err != nil {log.Printf("Error Occur while Converting string %d to %T", id, id)}
	product := dao.GetProduct(id)
	log.Printf("%+v", product)
	//TODO: Look at refactoring the json.Marshal & Unmarshal [DRY]
	response, err := json.Marshal(product)
	if err != nil {panic(err)}
	err = json.Unmarshal(response, &product)
	if err != nil {panic(err)}
	c.JSON(http.StatusOK, gin.H{"statusCode" : http.StatusOK, "response":product})
}


// CreateProductCategory POST /

func CreateProductCategoryCtrl(c *gin.Context) {
	//  Let gin do the parsing of the body, by using a binding.
	var productCategory models.ProductCategory
	_ = c.BindJSON(&productCategory)
	success := dao.CreateProductCategory(productCategory) // Inserts the productCategory to the DB
	if !success {
		c.JSON(http.StatusCreated, gin.H{"response": "OK"})
		return
	}
	//c.JSON(http.StatusOK, c)
}

func GetProductsCategoriesCtl(c *gin.Context)  {
	productsCategories := dao.GetProductsCategories()
	//TODO: Look at refactoring the json.Marshal & Unmarshal [DRY]
	response, err := json.Marshal(productsCategories)
	if err != nil {panic(err)}
	err = json.Unmarshal(response, &productsCategories)
	if err != nil {panic(err)}
	c.JSON(http.StatusOK, gin.H{"statusCode" : http.StatusOK, "response": productsCategories})
}

func GetProductCategoryCtrl(c *gin.Context){
	productCategoryId := c.Param("id")
	id, err := strconv.Atoi(productCategoryId)
	if err != nil {log.Printf("Error Occur while Converting string %d to %T", id, id)}
	productCategory := dao.GetProductCategory(id)
	//TODO: Look at refactoring the json.Marshal & Unmarshal [DRY]
	response, err := json.Marshal(productCategory)
	if err != nil {panic(err)}
	err = json.Unmarshal(response, &productCategory)
	if err != nil {panic(err)}
	c.JSON(http.StatusOK, gin.H{"statusCode" : http.StatusOK, "response": productCategory})
}