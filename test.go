package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var dbByAddress = make(map[string]string)
var dbByContract = make(map[string]string)
var hotByContract = make(map[string]int)
var coldByContract = make(map[string]int)

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/address/:address", func(c *gin.Context) {
		address := c.Params.ByName("address")
		value, ok := dbByAddress[address]

		if ok {
			c.JSON(http.StatusOK, gin.H{"address": address, "contract": value})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"address": address, "contract": "no-value"})
		}
	})

	r.GET("/contract/:contract", func(c *gin.Context) {
		contract := c.Params.ByName("contract")
		value, ok := dbByContract[contract]

		if ok {
			hot := hotByContract[contract]
			cold := coldByContract[contract]
			c.JSON(http.StatusOK, gin.H{"contract": contract, "address": value, "hot": hot, "cold": cold})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"contract": contract, "address": "no-value"})
		}
	})

	return r

}

func main() {

	//Быдлокод
	dbByAddress["Партизанская 40"] = "12345"
	dbByAddress["Комсомольский 10"] = "34567"

	dbByContract["34567"] = "Комсомольский 10"
	dbByContract["12345"] = "Партизанская 40"

	hotByContract["12345"] = 10
	hotByContract["34567"] = 20

	coldByContract["12345"] = 35
	coldByContract["34567"] = 32

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")

}
