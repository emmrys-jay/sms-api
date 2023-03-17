package router

import (
	"fmt"
	"github.com/Emmrys-Jay/altschool-sms/pkg/handler/student"
	"github.com/Emmrys-Jay/altschool-sms/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Student(r *gin.Engine, validate *validator.Validate, ApiVersion string, logger *utility.Logger) *gin.Engine {

	student := student.Controller{Validate: validate, Logger: logger}

	authUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		authUrl.POST("/student", student.Create)
		authUrl.GET("/student/:id", student.Get)
		authUrl.GET("/student", student.List)
		authUrl.PUT("/student/:id", student.Update)
		authUrl.DELETE("/student/:id", student.Delete)
		authUrl.GET("/student/:id/courses", student.ListStudentCourses)
		authUrl.PUT("/student/:id/courses", student.UpdateCourses)
	}
	return r
}
