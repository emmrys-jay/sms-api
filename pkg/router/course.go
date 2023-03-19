package router

import (
	"fmt"
	"github.com/Emmrys-Jay/altschool-sms/pkg/handler/course"
	"github.com/Emmrys-Jay/altschool-sms/pkg/middleware"
	"github.com/Emmrys-Jay/altschool-sms/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Course registers course handlers
func Course(r *gin.Engine, validate *validator.Validate, ApiVersion string, logger *utility.Logger) *gin.Engine {

	courseCtrl := course.Controller{Validate: validate, Logger: logger}

	authUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion), middleware.Admin())
	{
		authUrl.POST("/course", courseCtrl.Create)
		authUrl.GET("/course/:code", courseCtrl.Get)
		authUrl.GET("/course", courseCtrl.List)
		authUrl.PUT("/course/:code", courseCtrl.Update)
		authUrl.DELETE("/course/:code", courseCtrl.Delete)
		authUrl.GET("/course/:code/students", courseCtrl.ListCourseStudents)
	}

	return r
}
