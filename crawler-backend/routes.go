package main

import (
	"github.com/gin-gonic/gin"
)

// LookupNumber POST for frontend data lookup request
func lookupNumber(c *gin.Context) {
	// Allow CORS here By * or specific origin
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")

	// Bind the body to the Person struct
	data := &person{}
	c.Bind(data)

	// Create return object
	response := queriedPerson{" ", " ", -2}

	// Check if the number already exists
	// If so, parse the DB
	// else, crawl the site
	if checkIfStored(data.Number) {
		response = findMatchFromDB(data)
	} else {
		response = crawlSite(data)
	}

	c.JSON(200, gin.H{
		"status":  "posted",
		"message": response,
	})
}

// LookupNumberOption Discard the OPTIONS message due to CORS
func lookupNumberOption(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.AbortWithStatus(200)
}

// initializeRoutes defines all routes for API
func initializeRoutes() {
	// Define the endpoints
	router.POST("/lookup", lookupNumber)
	router.OPTIONS("/lookup", lookupNumberOption)
}
