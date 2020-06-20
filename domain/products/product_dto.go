package products

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ProductDetails struct {
	ProductId         int64     `gorm:"primary_key ;column:id" json:"productId"`
	Quantity          int       `gorm:"column:quantity" json:"quantity"`
	CreatedAt         time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt         time.Time `gorm:"column:updated_at" json:"-"`
	AvailableQuantity int       `gorm:"column:available_quantity" json:"availableQuantity"`
	Fare              float64   `gorm:"column:fare" json:"fare"`
	DiscountAmt       float64   `gorm:"column:discount_amt" json:"discountAmt"`
}

// set Segment table name to be `segments`
func (ProductDetails) TableName() string {
	return "product_details"
}

func (productDetail *ProductDetails) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("created_at", time.Now().UTC())
	scope.SetColumn("updated_at", time.Now().UTC())
	return nil
}

func (productDetail *ProductDetails) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("updated_at", time.Now().UTC())
	return nil
}
