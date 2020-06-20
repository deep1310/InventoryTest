package products

import "inventory/utils/errors"

type ProductQtyUpdate struct {
	ProductId  int64 `json:"productId"`
	Quantity   int   `json:"quantity"`
	IsRollback bool  `json:"isRollback"`
}

func (r *ProductQtyUpdate) ValidateProductQtyUpdateReq() *errors.RestErr {

	if r.ProductId <= 0 {
		return errors.BadRequestError("Not a valid request")
	}

	if r.Quantity <= 0 {
		return errors.BadRequestError("Not a valid request")
	}
	return nil
}
