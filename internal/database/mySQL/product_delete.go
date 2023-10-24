package mysql

import "github.com/adhimaswaskita/go-product-management/internal/models"

func (ms *MySQL) DeleteProduct(id uint) error {
	res := ms.DB.Delete(&models.Product{}, id)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
