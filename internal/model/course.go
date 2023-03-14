package model

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	CourseCode string `gorm:"column:course_code;" validate:"alphanum"`
	Units      int    `gorm:"column:units;"`
	Lecturer   string `gorm:"column:lecturer;"`
}
