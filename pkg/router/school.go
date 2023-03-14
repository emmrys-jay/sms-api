package router

import (
	"github.com/Emmrys-Jay/altschool-sms/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Student(r *gin.Engine, validate *validator.Validate, ApiVersion string, logger *utility.Logger) *gin.Engine {

	//health := health.Controller{Validate: validate, Logger: logger}
	//
	//authUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion))
	//{
	//	authUrl.POST("/student", health.Post)
	//	authUrl.GET("/student/:id", health.Get)
	//	authUrl.PUT("/student/:id", health.Update)
	//	authUrl.DELETE("/student/:id", health.Delete)
	//	authUrl.GET("/student/:id/courses", health.GetCourses)
	//	authUrl.PATCH("/student/:id/courses", health.UpdateCourses)
	//}
	return r
}
