package mysql

import "github.com/adhimaswaskita/go-product-management/internal/models"

func (ms *MySQL) DeleteAdmin(id uint) error {
	res := ms.DB.Delete(&models.Admin{}, id)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
