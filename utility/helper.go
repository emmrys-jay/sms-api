package utility

import (
	"github.com/Emmrys-Jay/altschool-sms/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func PasswordIsValid(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func SyncStudent(std *model.CreateStudent, student *model.Student) {
	student.FirstName = std.FirstName
	student.LastName = std.LastName
	student.Email = std.Email
	student.MatricNumber = std.MatricNumber
	student.YearOfBirth = std.YearOfBirth
	student.Courses = std.Courses
}
