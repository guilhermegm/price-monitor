package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	products "price-monitor/products"
)


func main() {
	session, err := mgo.Dial("104.131.170.117")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	router := gin.Default()

	products.CreateApi(session, router)

	router.Run(":3100")
}
