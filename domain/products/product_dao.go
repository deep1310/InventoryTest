package products

import (
	"inventory/datarepository/mysql/inventory_db"
	"inventory/utils/errors"
)

func (productDetail *ProductDetails) Get() *errors.RestErr {
	db := inventory_db.GetSqlConn()
	err := db.Where("id = ?", productDetail.ProductId).Find(&productDetail).Error
	if err != nil {
		return errors.InternalServerError("unable to get the product details")
	}
	return nil
}

func (r *ProductQtyUpdate) UpdateProductQuantity() *errors.RestErr {
	db := inventory_db.GetSqlConn()

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	prodDetails := &ProductDetails{}
	err := tx.Where("id = ?", r.ProductId).Find(&prodDetails).Error
	if err != nil {
		return errors.InternalServerError("unable to get the product details")
	}

	if !r.IsRollback {
		totalRowsUpdated := tx.Model(prodDetails).Where("id = ? AND available_quantity >= ?", r.ProductId, r.Quantity).Update(map[string]interface{}{"quantity": prodDetails.AvailableQuantity - r.Quantity}).RowsAffected
		if totalRowsUpdated == 0 {
			return errors.InternalServerError("unable to update product details")
		}
	} else {
		totalRowsUpdated := tx.Model(prodDetails).Where("id = ?", r.ProductId).Update(map[string]interface{}{"quantity": prodDetails.AvailableQuantity + r.Quantity}).RowsAffected
		if totalRowsUpdated == 0 {
			return errors.InternalServerError("unable to update product details")
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.InternalServerError("unable to update product details")
	}
	return nil
}
