package models

type ProductCategory struct {
	ID          uint   `gorm:"autoIncrement" json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
