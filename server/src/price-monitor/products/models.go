package products

import (
	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	Id string
	Name string
	Price float64
}

type Products []Product

type Price struct {
	ID bson.ObjectId
	Price float64
}

type Prices []Price

type ProductModel struct {
	productRepository ProductRepository
	priceRepository PriceRepository
}

func (productModel ProductModel) Find(query string) Products {
	queryBson := bson.M{"name": bson.M{"$regex": bson.RegEx{query, "i"}}}
	products := productModel.productRepository.Find(queryBson)
	return products
}
