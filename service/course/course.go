package course

import (
	"context"
	"github.com/Emmrys-Jay/altschool-sms/internal/model"
	"github.com/Emmrys-Jay/altschool-sms/pkg/repository/storage/mysql"
)

func Create(ctx context.Context, course *model.Course) error {

	db := mysql.GetDB()
	return db.CreateCourse(ctx, course)
}

func Get(ctx context.Context, code string) (*model.Course, error) {

	db := mysql.GetDB()
	return db.GetCourse(ctx, code)
}

func List(ctx context.Context) ([]model.Course, error) {

	db := mysql.GetDB()
	return db.ListCourses(ctx)
}

func ListStudents(ctx context.Context, code string) (*model.ListStudentsForm, error) {

	db := mysql.GetDB()
	return db.ListCourseStudents(ctx, code)
}

func Update(ctx context.Context, code string, course *model.Course) error {

	db := mysql.GetDB()
	return db.UpdateCourse(ctx, code, course)
}

func Delete(ctx context.Context, code string) error {

	db := mysql.GetDB()
	return db.DeleteCourse(ctx, code)
}
