package products

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProductRepository struct {
	productsCollection *mgo.Collection
}

type PriceRepository struct {
	pricesCollection *mgo.Collection
}

func (productsCollection ProductRepository) Find(query bson.M) Products {
	results := Products{}

	productsCollection.productsCollection.Find(query).All(&results)
	/* if ; err = nil {

	} */

	return results
}

func (priceRepository PriceRepository) Find(query bson.M) Prices {
	results := Prices{}

	priceRepository.pricesCollection.Find(query).All(&results)

	return results
}
