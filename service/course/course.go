package course

import (
	"context"
	"fmt"
	"github.com/Emmrys-Jay/altschool-sms/internal/model"
	"github.com/Emmrys-Jay/altschool-sms/pkg/repository/storage/mysql"
)

// Create creates a course from the details in 'course'
func Create(ctx context.Context, course *model.Course) error {

	db := mysql.GetDB()
	err := db.CreateCourse(ctx, course)
	if err != nil {
		return fmt.Errorf("could not create course, got error: %w", err)
	}

	return nil
}

// Get returns a course with 'code'
func Get(ctx context.Context, code string) (*model.Course, error) {

	db := mysql.GetDB()
	course, err := db.GetCourse(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("could not get course, got error: %w", err)
	}

	return course, nil
}

// List returns the total number of recognized courses
func List(ctx context.Context) ([]model.Course, error) {

	db := mysql.GetDB()
	courses, err := db.ListCourses(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not list courses, got error: %w", err)
	}

	return courses, nil
}

// ListStudents returns the students currently taking the course with 'code'
func ListStudents(ctx context.Context, code string) (*model.ListStudentsForm, error) {

	db := mysql.GetDB()
	stds, err := db.ListCourseStudents(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("could not list students, got error: %w", err)
	}

	return stds, nil
}

// Update updates a course with 'code' using the details in 'course'
func Update(ctx context.Context, code string, crs *model.UpdateCourse) error {
	course := model.Course{
		Title:    crs.Title,
		Lecturer: crs.Lecturer,
		Units:    crs.Units,
	}

	db := mysql.GetDB()
	err := db.UpdateCourse(ctx, code, &course)
	if err != nil {
		return fmt.Errorf("could not update course, got error: %w", err)
	}

	return nil
}

func Delete(ctx context.Context, code string) error {

	db := mysql.GetDB()
	err := db.DeleteCourse(ctx, code)
	if err != nil {
		return fmt.Errorf("could not delete course, got error: %w", err)
	}

	return nil
}
