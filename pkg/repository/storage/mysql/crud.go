package mysql

import (
	"context"
	"github.com/Emmrys-Jay/altschool-sms/internal/model"
)

// Course

func (m *MySQL) CreateCourse(ctx context.Context, course *model.Course) error {
	db, cancel := m.DBWithTimeout(ctx)
	defer cancel()
	return db.Create(course).Error
}

func (m *MySQL) GetCourse(ctx context.Context, id uint) (*model.Course, error) {
	db, cancel := m.DBWithTimeout(ctx)
	defer cancel()

	var course model.Course
	if err := db.First(&course, "id = ?", id).Error; err != nil {
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

func (m *MySQL) UpdateCourse(ctx context.Context, id uint, course *model.Course) error {
	db, cancel := m.DBWithTimeout(ctx)
	defer cancel()

	return db.Model(model.Course{}).Where("id = ?", id).Save(course).Error
}

func (m *MySQL) DeleteCourse(ctx context.Context, id uint) error {
	db, cancel := m.DBWithTimeout(ctx)
	defer cancel()

	return db.Where("id = ?", id).Delete(&model.Course{}).Error
}

// Student

func (m *MySQL) CreateStudent(ctx context.Context, student *model.Student) error {
	db, cancel := m.DBWithTimeout(ctx)
	defer cancel()
	return db.Create(student).Error
}

func (m *MySQL) GetStudent(ctx context.Context, id uint) (*model.Student, error) {
	db, cancel := m.DBWithTimeout(ctx)
	defer cancel()

	var student model.Student
	if err := db.First(&student, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &student, nil
}

func (m *MySQL) ListStudents(ctx context.Context) ([]model.Student, error) {
	db, cancel := m.DBWithTimeout(ctx)
	defer cancel()

	var students []model.Student
	if err := db.Find(&students).Error; err != nil {
		return nil, err
	}

	return students, nil
}

func (m *MySQL) UpdateStudent(ctx context.Context, id uint, student *model.Student) error {
	db, cancel := m.DBWithTimeout(ctx)
	defer cancel()

	return db.Model(model.Student{}).Where("id = ?", id).Save(student).Error
}

func (m *MySQL) DeleteStudent(ctx context.Context, id uint) error {
	db, cancel := m.DBWithTimeout(ctx)
	defer cancel()

	return db.Where("id = ?", id).Delete(&model.Student{}).Error
}
