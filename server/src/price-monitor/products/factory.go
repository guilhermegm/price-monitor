package products

import (
	"gopkg.in/mgo.v2"
	"github.com/gin-gonic/gin"
)


type ProductFactory struct {
}

var ginRouter *gin.Engine
var mongoSession *mgo.Session
var productsCollection *mgo.Collection
var pricesCollection *mgo.Collection

func CreateApi(session *mgo.Session, router *gin.Engine) {
	ginRouter = router
	mongoSession = session
	productsCollection = mongoSession.DB("price_monitor").C("products")
	pricesCollection = mongoSession.DB("price_monitor").C("products_price")
	
	priceRepository := PriceRepository{pricesCollection}
	productRepository := ProductRepository{productsCollection}
	productModel := ProductModel{productRepository, priceRepository}
	productApi := ProductApi{"/v1/products", productModel}

	productApi.createRoutes()
}
