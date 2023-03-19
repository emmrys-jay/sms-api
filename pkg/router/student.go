package router

import (
	"fmt"
	"github.com/Emmrys-Jay/altschool-sms/pkg/handler/student"
	"github.com/Emmrys-Jay/altschool-sms/pkg/middleware"
	"github.com/Emmrys-Jay/altschool-sms/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Student registers all student handlers
func Student(r *gin.Engine, validate *validator.Validate, ApiVersion string, logger *utility.Logger) *gin.Engine {

	studentCtrl := student.Controller{Validate: validate, Logger: logger}

	authUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		// Admin only
		authUrl.POST("/student", middleware.Admin(), studentCtrl.Create)
		authUrl.DELETE("/student/:id", middleware.Admin(), studentCtrl.Delete)
		authUrl.PUT("/student/:id/courses", middleware.Admin(), studentCtrl.UpdateCourses)
		authUrl.GET("/student", middleware.Admin(), studentCtrl.List)

		authUrl.GET("/student/:id", middleware.Student(), studentCtrl.Get)
		authUrl.PUT("/student/:id", middleware.Student(), studentCtrl.Update)
		authUrl.GET("/student/:id/courses", middleware.Student(), studentCtrl.ListStudentCourses)

	}

	return r
}
