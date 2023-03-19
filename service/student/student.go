package student

import (
	"context"
	"fmt"
	"github.com/Emmrys-Jay/altschool-sms/internal/model"
	"github.com/Emmrys-Jay/altschool-sms/pkg/repository/storage/mysql"
	"github.com/Emmrys-Jay/altschool-sms/utility"
)

// Create creates a student using the std param
func Create(ctx context.Context, std *model.CreateStudent) (*model.Student, error) {

	hash, err := utility.HashPassword(std.Password)
	if err != nil {
		return nil, fmt.Errorf("could not hash user password, got error: %w", err)
	}

	student := model.Student{
		FirstName:    std.FirstName,
		LastName:     std.LastName,
		Email:        std.Email,
		MatricNumber: std.MatricNumber,
		YearOfBirth:  std.YearOfBirth,
	}
	student.Password = hash

	db := mysql.GetDB()
	err = db.CreateStudent(ctx, &student)
	if err != nil {
		return nil, fmt.Errorf("could not create student, got error: %w", err)
	}

	return &student, nil
}

// Get returns a student with id
func Get(ctx context.Context, id uint) (*model.Student, error) {

	db := mysql.GetDB()
	std, err := db.GetStudent(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("could not list students, got error %w", err)
	}

	return std, nil
}

// List returns all recognized students
func List(ctx context.Context) ([]model.Student, error) {

	db := mysql.GetDB()
	stds, err := db.ListStudents(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not list students, got error %w", err)
	}

	return stds, nil
}

// ListCourses returns all courses curently being taken by student with id
func ListCourses(ctx context.Context, id uint) (*model.ListCoursesForm, error) {

	db := mysql.GetDB()
	student, err := db.ListStudentCourses(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("could not list courses by a student, got error: %w", err)
	}

	if student.Courses == nil {
		student.Courses = []model.Course{}
	}

	return student, nil
}

// Update updates a student with 'id' with the details in 'student'
func Update(ctx context.Context, id uint, std *model.UpdateStudent) error {
	student := model.Student{
		FirstName:    std.FirstName,
		LastName:     std.LastName,
		Email:        std.Email,
		MatricNumber: std.MatricNumber,
		YearOfBirth:  std.YearOfBirth,
	}

	db := mysql.GetDB()
	err := db.UpdateStudent(ctx, id, &student)
	if err != nil {
		return fmt.Errorf("could not update student, got error %w", err)
	}

	return nil
}

// UpdateCourses updates the courses of student with 'id' with 'courses'
func UpdateCourses(ctx context.Context, id uint, courses []model.Course) error {

	db := mysql.GetDB()
	err := db.UpdateStudentCourses(ctx, id, courses)
	if err != nil {
		return fmt.Errorf("could not update courses, got error %w", err)
	}

	return nil
}

// Delete deletes a student with 'id'
func Delete(ctx context.Context, id uint) error {

	db := mysql.GetDB()
	err := db.DeleteStudent(ctx, id)
	if err != nil {
		return fmt.Errorf("could not delete student, got error %w", err)
	}

	return nil
}
