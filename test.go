package main

import (
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
)

var dbByAddress = make(map[string]string)
var dbByContract = make(map[string]string)

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/address/:address", func(c *gin.Context) {
		address := c.Params.ByName("address")
		value, ok := dbByAddress[address]

		fmt.Print(value)

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
			c.JSON(http.StatusOK, gin.H{"contract": contract, "address": value})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"contract": contract, "address": "no-value"})
		}
	})

	return r

}

func main() {

	//Быдлокод
	dbByAddress["Партизанская 40"] = "10000"
	dbByAddress["Комсомольский 10"] = "20000"

	dbByContract["20000"] = "Комсомольский 10"
	dbByContract["10000"] = "Партизанская 40"

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")

}
