package mysql

import "github.com/adhimaswaskita/go-product-management/internal/models"

func (ms *MySQL) CreateAdmin(admin models.Admin) error {
	res := ms.DB.Create(&admin)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
