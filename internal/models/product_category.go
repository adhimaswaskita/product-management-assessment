package models

type ProductCategory struct {
	ID          uint   `gorm:"autoIncrement" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
