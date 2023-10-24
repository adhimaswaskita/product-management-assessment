package models

type TransactionHistory struct {
	ID           uint `gorm:"autoIncrement"`
	Event        string
	ActivityType string
}
