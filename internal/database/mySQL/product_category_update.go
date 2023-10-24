package mysql

import "github.com/adhimaswaskita/go-product-management/internal/models"

func (ms *MySQL) UpdateProductCategory(id uint, param models.ProductCategory) error {
	param.ID = id
	res := ms.DB.Save(&param)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
