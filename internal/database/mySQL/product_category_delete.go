package mysql

import "github.com/adhimaswaskita/go-product-management/internal/models"

func (ms *MySQL) DeleteProductCategory(id uint) error {
	res := ms.DB.Delete(&models.ProductCategory{}, id)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
