package mysql

import "github.com/adhimaswaskita/go-product-management/internal/models"

func (ms *MySQL) GetProduct() ([]models.Product, error) {
	var products []models.Product
	rows, err := ms.DB.Model(&models.Product{}).
		Select(
			"products.id, products.name, products.description, products.image, products.stock, products.category_id, product_categories.id, product_categories.name, product_categories.description",
		).
		Joins("left join product_categories on products.category_id = product_categories.id").
		Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		product := &models.Product{}
		err := rows.Scan(
			&product.ID, &product.Name, &product.Description, &product.Image, &product.Stock, &product.CategoryID, &product.Category.ID, &product.Category.Name, &product.Category.Description,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
	}

	return products, nil
}
