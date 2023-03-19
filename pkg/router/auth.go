package router

import (
	"fmt"
	"github.com/Emmrys-Jay/altschool-sms/pkg/handler/user"
	"github.com/Emmrys-Jay/altschool-sms/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Auth registers authentication handlers
func Auth(r *gin.Engine, validate *validator.Validate, ApiVersion string, logger *utility.Logger) *gin.Engine {

	auth := user.Controller{Validate: validate, Logger: logger}

	authUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		authUrl.POST("/login", auth.Login)
	}
	return r
}
