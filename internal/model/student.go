package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	FirstName    string `gorm:"column:first_name;"`
	LastName     string `gorm:"column:last_name;"`
	MatricNumber string `gorm:"column:matric_number;" validate:"alphanum"`
	YearOfBirth  int    `gorm:"column:year_of_birth;" validate:"gte=0"`
}
