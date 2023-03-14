package router

import (
	"fmt"
	"github.com/Emmrys-Jay/altschool-sms/pkg/handler/course"
	"github.com/Emmrys-Jay/altschool-sms/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Course(r *gin.Engine, validate *validator.Validate, ApiVersion string, logger *utility.Logger) *gin.Engine {

	course := course.Controller{Validate: validate, Logger: logger}

	authUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		authUrl.POST("/courses", course.Create)
		authUrl.GET("/course/:id", course.Get)
		authUrl.GET("/course", course.List)
		authUrl.PUT("/course/:id", course.Update)
		authUrl.DELETE("/course/:id", course.Delete)
		//authUrl.GET("/course/:id/students", course.GetStudents)
	}

	return r
}
