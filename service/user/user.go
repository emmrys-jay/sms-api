package user

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"

	"github.com/Emmrys-Jay/altschool-sms/internal/config"
	"github.com/Emmrys-Jay/altschool-sms/pkg/middleware"
	"gorm.io/gorm"

	"github.com/Emmrys-Jay/altschool-sms/pkg/repository/storage/mysql"
	"github.com/Emmrys-Jay/altschool-sms/utility"

	"github.com/Emmrys-Jay/altschool-sms/internal/model"
)

// Login generates and returns a token for a recognized user
func Login(ctx context.Context, user *model.Login) (*model.LoginResponse, error) {

	db := mysql.GetDB()
	std, err := db.GetStudentByMatNum(ctx, user.MatNum)
	if err != nil {
		return nil, fmt.Errorf("couldn't get student, got error %w", err)
	}

	if !utility.PasswordIsValid(user.Password, std.Password) {
		return nil, errors.New("invalid matric number or password")
	}

	token, err := middleware.CreateToken(std.MatricNumber)
	if err != nil {
		return nil, fmt.Errorf("couldn't create token, got error %w", err)
	}

	logResponse := model.LoginResponse{
		User: model.ListCoursesForm{
			ID:           std.ID,
			MatricNumber: std.MatricNumber,
			FirstName:    std.FirstName,
			LastName:     std.LastName,
		},
		Token: token,
	}

	return &logResponse, nil
}

// CreateAdminUser creates an admin user if one doesn't exist
func CreateAdminUser(logger *utility.Logger) error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	matNum := config.GetConfig().AdminMatNum
	password, err := utility.HashPassword(config.GetConfig().AdminPassword)
	if err != nil {
		return fmt.Errorf("could not hash admin password, got error: %w", err)
	}

	db := mysql.GetDB()

	// Check if admin user already exists
	if _, err = db.GetStudentByMatNum(ctx, matNum); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Create admin user if it doesn't exist
			err := db.CreateStudent(ctx, &model.Student{
				FirstName:    matNum,
				LastName:     matNum,
				MatricNumber: matNum,
				Password:     password,
				Email:        matNum,
			})
			if err != nil {
				return fmt.Errorf("could not create admin user, got error: %w", err)
			}
		} else {
			return fmt.Errorf("could not check admin user, got error: %w", err)
		}
	} else {
		log.Info("ADMIN USER ALREADY EXISTS")
		logger.Info("ADMIN USER ALREADY EXISTS")
		return nil
	}

	log.Info("CREATED ADMIN USER SUCCESSFULLY")
	logger.Info("CREATED ADMIN USER SUCCESSFULLY")
	return nil
}
