package router

import (
	"github.com/Emmrys-Jay/altschool-sms/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Course(r *gin.Engine, validate *validator.Validate, ApiVersion string, logger *utility.Logger) *gin.Engine {

	//health := health.Controller{Validate: validate, Logger: logger}
	//
	//authUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion))
	//{
	//	authUrl.POST("/courses", health.Post)
	//	authUrl.GET("/course/:id", health.Get)
	//	authUrl.PUT("/course/:id", health.Update)
	//	authUrl.DELETE("/course/:id", health.Delete)
	//	authUrl.GET("/course/:id/students", health.GetStudents)
	//}
	return r
}
