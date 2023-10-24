package models

type Admin struct {
	ID          uint   `gorm:"autoIncrement"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name"`
	Email       string `json:"email" validate:"email"`
	DateOfBirth string `json:"date_of_birth"`
	Sex         string `json:"sex"`
	Password    string `json:"password"`
}
