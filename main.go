package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getCountry(c *gin.Context) {
	country := c.Param("country")
	result := GetCountry(country)

	if result != nil {
		c.IndentedJSON(http.StatusOK, result)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "country not found"})
}

func postCountry(c *gin.Context) {
	var newCountry country

	if err := c.BindJSON(&newCountry); err != nil {
		return
	}

	if countryAdded := AddCountry(newCountry); !countryAdded {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "country not avaible"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newCountry)
}

func updateCountry(c *gin.Context) {
	var update country

	if err := c.BindJSON(&update); err != nil {
		return
	}

	if countryDeleted := DeleteCountry(update.Country); !countryDeleted {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "country not found"})
		return
	}

	if countryAdded := AddCountry(update); !countryAdded {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "country not avaible"})
		return
	}

	c.IndentedJSON(http.StatusOK, update)
}

func deleteCountry(c *gin.Context) {
	country := c.Param("country")

	if countryDeleted := DeleteCountry(country); !countryDeleted {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "country not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "country deleted"})
}

func main() {
	router := gin.Default()

	router.GET("/countries/:country", getCountry)
	router.POST("/countries", postCountry)
	router.PUT("/countries", updateCountry)
	router.DELETE("countries/:country", deleteCountry)

	router.Run("localhost:8080")
}
