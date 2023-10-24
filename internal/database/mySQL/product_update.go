package mysql

import "github.com/adhimaswaskita/go-product-management/internal/models"

func (ms *MySQL) UpdateProduct(id uint, param models.Product) error {
	param.ID = id
	res := ms.DB.Save(&param)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
