package mysql

import "github.com/adhimaswaskita/go-product-management/internal/models"

func (ms *MySQL) GetAdmin() ([]models.Admin, error) {
	var admin []models.Admin
	res := ms.DB.Find(&admin)
	if res.Error != nil {
		return nil, res.Error
	}

	return admin, nil
}
