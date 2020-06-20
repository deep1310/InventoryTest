package app

import (
	"github.com/gin-gonic/gin"
	"inventory/controllers/inventory"
)

var router = gin.Default()

func Start() {
	router.GET("/ProductDetails/:product_id", inventory.GetProductDetails)
	router.POST("/ProductDetails/UpdateQuantity/", inventory.UpdateProductQuantity)
	router.Run(":5550")
}
