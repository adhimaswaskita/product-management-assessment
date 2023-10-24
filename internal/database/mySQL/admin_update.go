package mysql

import "github.com/adhimaswaskita/go-product-management/internal/models"

func (ms *MySQL) UpdateAdmin(id uint, param models.Admin) error {
	param.ID = id
	res := ms.DB.Save(&param)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
