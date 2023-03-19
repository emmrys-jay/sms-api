package model

import (
	"time"
)

type Course struct {
	CourseCode string    `gorm:"column:course_code;unique;primaryKey" validate:"required,alphanum" json:"course_code"`
	Title      string    `gorm:"column:title;" validate:"required" json:"title"`
	Units      int       `gorm:"column:units;" validate:"required" json:"units"`
	Lecturer   string    `gorm:"column:lecturer;" json:"lecturer"`
	CreatedAt  time.Time `gorm:"created_at;" json:"created_at"`
	UpdatedAt  time.Time `gorm:"updated_at;" json:"updated_at"`
	Students   []Student `gorm:"many2many:student_courses;" json:"students,omitempty"`
}

type CourseForDoc struct {
	CourseCode string `gorm:"column:course_code;unique;primaryKey" validate:"required,alphanum" json:"course_code"`
	Title      string `gorm:"column:title;" validate:"required" json:"title"`
	Units      int    `gorm:"column:units;" validate:"required" json:"units"`
	Lecturer   string `gorm:"column:lecturer;" json:"lecturer"`
}

type UpdateCourse struct {
	Title    string `gorm:"column:title;" json:"title"`
	Units    int    `gorm:"column:units;" json:"units"`
	Lecturer string `gorm:"column:lecturer;" json:"lecturer"`
}

type ListStudentsForm struct {
	CourseCode string    `gorm:"column:course_code;unique" validate:"required,alphanum" json:"course_code"`
	Title      string    `gorm:"column:title;" validate:"required" json:"title"`
	Units      int       `gorm:"column:units;" validate:"required" json:"units"`
	Lecturer   string    `gorm:"column:lecturer;" json:"lecturer"`
	Students   []Student `gorm:"many2many:student_courses;" json:"students,omitempty"`
}

func (Course) TableName() string {
	return "courses"
}
