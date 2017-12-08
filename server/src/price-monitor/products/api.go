package products

import (
	"github.com/gin-gonic/gin"
)

type ProductApi struct {
	urlPrefix string
	productModel ProductModel
}

func (productApi ProductApi) createGet() {
	ginRouter.GET(productApi.urlPrefix, func(context *gin.Context) {
		query := context.Query("query")
		products := productApi.productModel.Find(query)

		responseCode := 200
		response := gin.H{
			"items": products,
		}

		context.JSON(responseCode, response)
	})
}

func (productApi ProductApi) createRoutes() {
	productApi.createGet()
}
