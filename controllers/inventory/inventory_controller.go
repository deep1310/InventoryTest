package inventory

import (
	"github.com/gin-gonic/gin"
	"inventory/domain/products"
	"inventory/services"
	"inventory/utils/errors"
	"net/http"
	"strconv"
	"strings"
)

func GetProductDetails(c *gin.Context) {

	productId := c.Param("product_id")
	productId = strings.TrimSpace(productId)
	if productId == "" {
		apiReqErr := errors.BadRequestError("product id is empty")
		c.JSON(apiReqErr.Code, apiReqErr)
		return
	}

	productIdInt, userErr := strconv.ParseInt(productId, 10, 64)
	if userErr != nil {
		apiReqErr := errors.BadRequestError("product id is not int")
		c.JSON(apiReqErr.Code, apiReqErr)
		return
	}

	result, err := services.InventoryService.GetProductDetails(productIdInt)
	if err != nil {
		c.JSON(err.Code, result)
		return
	}
	c.JSON(http.StatusOK, result)
}

func UpdateProductQuantity(c *gin.Context) {

	var quantityUpdateReq products.ProductQtyUpdate
	if err := c.ShouldBindJSON(&quantityUpdateReq); err != nil {
		apiReqErr := errors.BadRequestError("invalid request")
		c.JSON(apiReqErr.Code, apiReqErr)
		return
	}

	err := services.InventoryService.UpdateProductQuantity(&quantityUpdateReq)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, "")

}
