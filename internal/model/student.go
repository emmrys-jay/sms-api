package model

import (
	"time"
)

type Student struct {
	ID           uint      `gorm:"column:id;type:uint;" json:"id"`
	FirstName    string    `gorm:"column:first_name;" json:"first_name"`
	LastName     string    `gorm:"column:last_name;" json:"last_name"`
	Email        string    `gorm:"column:email" json:"email"`
	MatricNumber string    `gorm:"column:matric_number;unique;not null;" json:"matric_number"`
	Password     string    `gorm:"column:password" json:"-"`
	YearOfBirth  int       `gorm:"column:year_of_birth;" json:"year_of_birth,omitempty"`
	CreatedAt    time.Time `gorm:"created_at;" json:"created_at"`
	UpdatedAt    time.Time `gorm:"updated_at;" json:"updated_at"`
	Courses      []Course  `gorm:"many2many:student_courses;" json:"courses,omitempty"`
}

type CreateStudent struct {
	FirstName    string `validate:"required" json:"first_name"`
	LastName     string `validate:"required" json:"last_name"`
	Email        string `validate:"email" json:"email"`
	MatricNumber string `validate:"required" json:"matric_number"`
	Password     string `validate:"required" json:"password"`
	YearOfBirth  int    `validate:"gte=0" json:"year_of_birth,omitempty"`
}

type UpdateStudent struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	MatricNumber string `json:"matric_number"`
	Password     string `json:"password"`
	YearOfBirth  int    `validate:"gte=0" json:"year_of_birth,omitempty"`
}

type ListCoursesForm struct {
	ID           uint     `gorm:"column:id;type:uint;" json:"id"`
	FirstName    string   `gorm:"column:first_name;" json:"first_name"`
	LastName     string   `gorm:"column:last_name;" json:"last_name"`
	MatricNumber string   `gorm:"column:matric_number;unique;not null;" json:"matric_number"`
	Courses      []Course `gorm:"many2many:student_courses;" json:"courses,omitempty"`
}

func (Student) TableName() string {
	return "students"
}
