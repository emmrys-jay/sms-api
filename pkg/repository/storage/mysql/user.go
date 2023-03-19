package mysql

import (
	"context"

	"github.com/Emmrys-Jay/altschool-sms/internal/model"
)

// GetStudentByMatNum return a student with matric number 'matNum'
func (m *MySQL) GetStudentByMatNum(ctx context.Context, matNum string) (*model.Student, error) {

	db, cancel := m.DBWithTimeout(ctx)
	defer cancel()

	var student model.Student
	err := db.First(&student, "matric_number = ?", matNum).Error
	if err != nil {
		return nil, err
	}

	return &student, nil
}
