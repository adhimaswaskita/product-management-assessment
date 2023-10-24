package models

type Product struct {
	ID          uint            `gorm:"autoIncrement" json:"id"`
	Name        string          `json:"name" validate:"required"`
	Description string          `json:"description"`
	Image       string          `json:"string"`
	CategoryID  int32           `json:"category_id" validate:"required"`
	Category    ProductCategory `gorm:"references:category_id" json:"category"`
	Stock       int             `json:"stock"`
}
