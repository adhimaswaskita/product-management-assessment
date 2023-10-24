package mysql

import "github.com/adhimaswaskita/go-product-management/internal/models"

func (ms *MySQL) CreateProductCategory(param models.ProductCategory) error {
	res := ms.DB.Create(&param)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
