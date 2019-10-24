package apigee

import "log"
import "fmt"

func GetUser(msisdn string) string {
	log.Println("[APIGEE SERVICES: SHOULD return User by MSISDN]")
	ApiGeeGetUsersUrl := fmt.Sprintf("GET: https://api.apigee/users/%s", msisdn)
	log.Println(ApiGeeGetUsersUrl)
	response := string(`{"MSISDN":"0731209834", "Name":"Jack Daniels", "Email":"jack.daniels@emil.com"}`)
	return response
}


func GetBundle(bundleId int)  string {
	log.Println("[APIGEE SERVICES: SHOULD return Bundle by bundleId]")
	ApiGeeGetBundleUrl := "GET: https://api.apigee/bundles/" + string(bundleId)
	log.Println(ApiGeeGetBundleUrl)
	response := string(`{"BundleID":1, "BundleTypeID":1, "BundleValidity":"30 Days", "BundleCost":30.00, "BundleValue":100}`)
	return response
}

func GetProduct(bundleTypeId int) string {
	log.Println("[APIGEE SERVICES: SHOULD ret	urn BundleType by bundleTypeID]")
	ApiGeeGetProductsUrl := fmt.Sprintf("GET: https://api.apigee/bundlestypes/%s", string(bundleTypeId))
	log.Println(ApiGeeGetProductsUrl)
	response := string(`{"BundleTypeID":1,"BundleTypeName":"Airtime", "BundleTypeDescription":"Airtime Bundle Type"}`)
	return response
}


func GetCategory(categoryId string)  {
	log.Println("[APIGEE SERVICES: SHOULD return Category by categoryId]")
	return
}

func GetProduct(productId int)  string {
	log.Println("[APIGEE SERVICES: SHOULD return Product by productId]")
	ApiGeeGetProductsUrl := fmt.Sprintf("GET: https://api.apigee/products/%s", string(productId))
	log.Println(ApiGeeGetProductsUrl)
	response := string(`{"ProductName": "Social Promo Data","ProductID":1, "Bundles":[{"BundleTypeID":1, "BundleID":1, "BundleValidity"":"24 Hours","BundleCost":10.00, "BundleValue":75}]}`)
	return response
}

func GetCombo(comboId int)  string {
	log.Println("[APIGEE SERVICES: SHOULD return Combo by comboId]")
	ApiGeeGetCombosUrl := fmt.Sprintf("GET: https://api.apigee/products/%s", string(comboId))
	log.Println(ApiGeeGetCombosUrl)
	response := string(`{"ComboID":1,"ComboCost":55,"ComboValidity": "1 Week"}`)
	return response
}