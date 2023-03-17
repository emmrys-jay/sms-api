package mysql

import (
	"context"
	"github.com/Emmrys-Jay/altschool-sms/internal/model"
	"gorm.io/gorm"
)

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

func (m *MySQL) ListStudentCourses(ctx context.Context, id uint) (*model.ListCoursesForm, error) {
	db, cancel := m.DBWithTimeout(ctx)
	defer cancel()

	var student model.Student
	err := db.Preload("Courses").Where("id = ?", id).First(&student).Error
	if err != nil {
		return nil, err
	}

	listCoursesForm := model.ListCoursesForm{
		ID:           student.ID,
		FirstName:    student.FirstName,
		LastName:     student.LastName,
		MatricNumber: student.MatricNumber,
		Courses:      student.Courses,
	}

	return &listCoursesForm, nil
}

func (m *MySQL) UpdateStudent(ctx context.Context, id uint, student *model.Student) error {
	db, cancel := m.DBWithTimeout(ctx)
	defer cancel()

	return db.Model(model.Student{}).Where("id = ?", id).Updates(student).Error
}

func (m *MySQL) UpdateStudentCourses(ctx context.Context, id uint, newCourses []model.Course) error {
	db, cancel := m.DBWithTimeout(ctx)
	defer cancel()

	var student model.Student
	// Get current student courses
	err := db.Preload("Courses").Where("id = ?", id).First(&student).Error
	if err != nil {
		return err
	}

	// If student already takes some courses, disassociate them from the courses
	if len(student.Courses) != 0 {
		err := db.Model(&student).Association("Courses").Delete(student.Courses)
		if err != nil {
			return err
		}
	}

	// Associate new courses
	student.Courses = newCourses
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&student).Error
}

func (m *MySQL) DeleteStudent(ctx context.Context, id uint) error {
	db, cancel := m.DBWithTimeout(ctx)
	defer cancel()

	return db.Where("id = ?", id).Delete(&model.Student{}).Error
}
