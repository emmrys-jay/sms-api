package student

import (
	"github.com/Emmrys-Jay/altschool-sms/utility"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
	Validate *validator.Validate
	Logger   *utility.Logger
}
