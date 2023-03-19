package router

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"

	"github.com/Emmrys-Jay/altschool-sms/pkg/middleware"
	"github.com/Emmrys-Jay/altschool-sms/utility"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	_ "github.com/Emmrys-Jay/altschool-sms/docs"
)

// Setup creates a new router and registers all handlers to it
func Setup(validate *validator.Validate, logger *utility.Logger) *gin.Engine {
	r := gin.New()

	// Middlewares
	// r.Use(gin.Logger())
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	ApiVersion := "v1"
	Health(r, validate, ApiVersion, logger)
	Auth(r, validate, ApiVersion, logger)
	Student(r, validate, ApiVersion, logger)
	Course(r, validate, ApiVersion, logger)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"name":    "Not Found",
			"message": "Page not found.",
			"code":    404,
			"status":  http.StatusNotFound,
		})
	})

	return r
}
