package mysql

func (ms *MySQL) UpdateProductStock(id uint, stock int) error {
	qs := "UPDATE products SET stock = "
	qs += "CASE WHEN ((? < 0) AND (stock + ?) < 0) THEN stock "
	qs += "ELSE (stock + ?) END "
	qs += "WHERE id = ?;"
	res := ms.DB.Exec(qs, stock, stock, stock, id)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
