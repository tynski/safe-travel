package main

import (
	"github.com/gin-gonic/gin"
)

func getCountry(c *gin.Context) {

}

func postCountry(c *gin.Context) {

}

func updateCountry(c *gin.Context) {

}

func deleteCountry(c *gin.Context) {

}

func main() {
	router := gin.Default()

	router.GET("/countries/:country", getCountry)
	router.POST("/countries", postCountry)
	router.PUT("/countries/:country", updateCountry)
	router.DELETE("countries/:country", deleteCountry)

	router.Run("localhost:8080")
}
