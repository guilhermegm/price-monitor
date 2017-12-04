package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	ID     bson.ObjectId
	Name  string
 }

type Products []Product

func main() {
	session, err := mgo.Dial("104.131.170.117")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	productsCollection := session.DB("price_monitor").C("products")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		var responseCode int
		var response gin.H
		query := c.Query("query")
		results := Products{}
		queryBson := bson.M{"name": bson.M{"$regex": bson.RegEx{query, "i"}}}
		
		if err := productsCollection.Find(queryBson).All(&results); err != nil {
			fmt.Println("Failed to write results:", err)
			responseCode = 404
			response = gin.H{
				"message": "No data",
			}
		} else {
			responseCode = 200
			response = gin.H{
				"data": results,
			}
		}
		fmt.Println("Results:", results)
	

		c.JSON(responseCode, response)
	})
	r.Run(":3100")
}
