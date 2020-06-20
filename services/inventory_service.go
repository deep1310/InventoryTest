package services

import (
	"inventory/domain/products"
	"inventory/utils/errors"
)

var (
	InventoryService InventoryServiceInterface = &inventoryServiceRepo{}
)

type inventoryServiceRepo struct{}

type InventoryServiceInterface interface {
	GetProductDetails(int64) (*products.ProductDetails, *errors.RestErr)
	UpdateProductQuantity(*products.ProductQtyUpdate) *errors.RestErr
}

func (s *inventoryServiceRepo) GetProductDetails(productId int64) (*products.ProductDetails, *errors.RestErr) {
	/*
		Note : Here the offer will come from the offer microservice
		and will be mapped to product response.
		I have put the same as of now part of product_details table to
		show the usage
	*/

	productDetail := &products.ProductDetails{ProductId: productId}
	if err := productDetail.Get(); err != nil {
		return nil, err
	}
	return productDetail, nil
}

func (s *inventoryServiceRepo) UpdateProductQuantity(req *products.ProductQtyUpdate) *errors.RestErr {

	if err := req.ValidateProductQtyUpdateReq(); err != nil {
		return err
	}

	if err := req.UpdateProductQuantity(); err != nil {
		return err
	}
	return nil
}
