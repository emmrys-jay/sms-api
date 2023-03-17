package storage

import (
	"context"
	"github.com/Emmrys-Jay/altschool-sms/internal/model"
)

// Repository is the database interface
type Repository interface {
	// Course
	CreateCourse(ctx context.Context, course *model.Course) error
	GetCourse(ctx context.Context, code string) (*model.Course, error)
	ListCourses(ctx context.Context) ([]model.Course, error)
	UpdateCourse(ctx context.Context, code string, course *model.Course) error
	DeleteCourse(ctx context.Context, code string) error

	ListCourseStudents(ctx context.Context, code string) (*model.ListStudentsForm, error)

	// Student
	CreateStudent(ctx context.Context, student *model.Student) error
	GetStudent(ctx context.Context, id uint) (*model.Student, error)
	ListStudents(ctx context.Context) ([]model.Student, error)
	UpdateStudent(ctx context.Context, id uint, student *model.Student) error
	DeleteStudent(ctx context.Context, id uint) error

	ListStudentCourses(ctx context.Context, id uint) (*model.ListCoursesForm, error)
	UpdateStudentCourses(ctx context.Context, id uint, courses []model.Course) error
}

// repositories
