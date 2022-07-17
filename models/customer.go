package models

type Customer struct {
	ID          int    `gorm:"column:id" json:"id"`
	Name        string `gorm:"column:name" validate:"required"`
	PhoneNumber string `gorm:"column:phone_number" validate:"required"`
}
