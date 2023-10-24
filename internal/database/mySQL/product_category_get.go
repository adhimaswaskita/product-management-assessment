package mysql

import "github.com/adhimaswaskita/go-product-management/internal/models"

func (ms *MySQL) GetProductCategory() ([]models.ProductCategory, error) {
	var admin []models.ProductCategory
	res := ms.DB.Find(&admin)
	if res.Error != nil {
		return nil, res.Error
	}

	return admin, nil
}
