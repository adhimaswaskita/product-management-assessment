package mysql

import "github.com/adhimaswaskita/go-product-management/internal/models"

func (ms *MySQL) GetProduct() ([]models.Product, error) {
	var products []models.Product
	rows, err := ms.DB.Model(&models.Product{}).
		Select("product.id, product.name, product.description, product_categories.id, product_categories.name").
		Joins("left join product_categories on products.category_id = product_categories.id").
		Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		product := &models.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Category.ID, &product.Category.Name)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
	}

	return products, nil
}
