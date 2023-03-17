package student

import (
	"context"
	"github.com/Emmrys-Jay/altschool-sms/internal/model"
	"github.com/Emmrys-Jay/altschool-sms/pkg/repository/storage/mysql"
	"github.com/Emmrys-Jay/altschool-sms/utility"
)

func Create(ctx context.Context, std *model.CreateStudent) (*model.Student, error) {

	hash, err := utility.HashPassword(std.Password)
	if err != nil {
		return nil, err
	}

	var student model.Student
	utility.SyncStudent(std, &student)
	student.Password = hash

	db := mysql.GetDB()
	err = db.CreateStudent(ctx, &student)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func Get(ctx context.Context, id uint) (*model.Student, error) {

	db := mysql.GetDB()
	return db.GetStudent(ctx, id)
}

func List(ctx context.Context) ([]model.Student, error) {

	db := mysql.GetDB()
	return db.ListStudents(ctx)
}

func ListCourses(ctx context.Context, id uint) (*model.ListCoursesForm, error) {

	db := mysql.GetDB()
	student, err := db.ListStudentCourses(ctx, id)
	if err != nil {
		return nil, err
	}

	if student.Courses == nil {
		student.Courses = []model.Course{}
	}

	return student, nil
}

func Update(ctx context.Context, id uint, student *model.Student) error {
	db := mysql.GetDB()
	return db.UpdateStudent(ctx, id, student)
}

func UpdateCourses(ctx context.Context, id uint, courses []model.Course) error {

	db := mysql.GetDB()
	return db.UpdateStudentCourses(ctx, id, courses)
}

func Delete(ctx context.Context, id uint) error {

	db := mysql.GetDB()
	return db.DeleteStudent(ctx, id)
}
