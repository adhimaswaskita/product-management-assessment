package models

type Product struct {
	ID          uint            `gorm:"autoIncrement" json:"id"`
	Name        string          `json:"name" validate:"required"`
	Description string          `json:"description"`
	Image       string          `json:"image"`
	CategoryID  uint            `json:"category_id" validate:"required"`
	Category    ProductCategory `gorm:"foreignKey:category_id" json:"category"`
	Stock       int             `json:"stock"`
}
