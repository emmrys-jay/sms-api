package router

import (
	"fmt"

	"github.com/Emmrys-Jay/altschool-sms/pkg/handler/health"
	"github.com/Emmrys-Jay/altschool-sms/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Health(r *gin.Engine, validate *validator.Validate, ApiVersion string, logger *utility.Logger) *gin.Engine {

	health := health.Controller{Validate: validate, Logger: logger}

	authUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		authUrl.POST("/health", health.Post)
		authUrl.GET("/health", health.Get)
	}
	return r
}
