package storage

import (
	"context"
	"github.com/Emmrys-Jay/altschool-sms/internal/model"
	"gorm.io/gorm"
)

// Repository is the database interface
type Repository interface {
	ConnectToDB() *gorm.DB

	// Course
	CreateCourse(ctx context.Context, course *model.Course) error
	GetCourse(ctx context.Context, id uint) (*model.Course, error)
	ListCourses(ctx context.Context) ([]model.Course, error)
	UpdateCourse(ctx context.Context, id uint, course *model.Course) error
	DeleteCourse(ctx context.Context, id uint) error

	// Student
	CreateStudent(ctx context.Context, course *model.Student) error
	GetStudent(ctx context.Context, id uint) (*model.Student, error)
	ListStudents(ctx context.Context) ([]model.Student, error)
	UpdateStudent(ctx context.Context, id uint, course *model.Student) error
	DeleteStudent(ctx context.Context, id uint) error
}

// repositories
