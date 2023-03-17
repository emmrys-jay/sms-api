package mysql

import (
	"context"
	"github.com/Emmrys-Jay/altschool-sms/internal/model"
)

func (m *MySQL) CreateCourse(ctx context.Context, course *model.Course) error {
	db, cancel := m.DBWithTimeout(ctx)
	defer cancel()
	return db.Create(course).Error
}

func (m *MySQL) GetCourse(ctx context.Context, code string) (*model.Course, error) {
	db, cancel := m.DBWithTimeout(ctx)
	defer cancel()

	var course model.Course
	if err := db.First(&course, "course_code = ?", code).Error; err != nil {
		return nil, err
	}

	return &course, nil
}

func (m *MySQL) ListCourses(ctx context.Context) ([]model.Course, error) {
	db, cancel := m.DBWithTimeout(ctx)
	defer cancel()

	var courses []model.Course
	if err := db.Find(&courses).Error; err != nil {
		return nil, err
	}

	return courses, nil
}

func (m *MySQL) ListCourseStudents(ctx context.Context, code string) (*model.ListStudentsForm, error) {
	db, cancel := m.DBWithTimeout(ctx)
	defer cancel()

	var course model.Course
	err := db.Preload("Students").Where("course_code = ?", code).First(&course).Error
	if err != nil {
		return nil, err
	}

	listStudentsForm := model.ListStudentsForm{
		CourseCode: course.CourseCode,
		Units:      course.Units,
		Title:      course.Title,
		Lecturer:   course.Lecturer,
		Students:   course.Students,
	}

	return &listStudentsForm, nil
}

func (m *MySQL) UpdateCourse(ctx context.Context, code string, course *model.Course) error {
	db, cancel := m.DBWithTimeout(ctx)
	defer cancel()

	return db.Model(model.Course{}).Where("course_code = ?", code).Updates(course).Error
}

func (m *MySQL) DeleteCourse(ctx context.Context, code string) error {
	db, cancel := m.DBWithTimeout(ctx)
	defer cancel()

	return db.Where("course_code = ?", code).Delete(&model.Course{}).Error
}
