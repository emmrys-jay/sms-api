package user

import (
	"net/http"

	"github.com/Emmrys-Jay/altschool-sms/utility"
	"github.com/gin-gonic/gin"
)

func (base *Controller) CreateUser(c *gin.Context) {

	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", gin.H{"user": "user object"})
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) Login(c *gin.Context) {

	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", gin.H{"user": "login object"})
	c.JSON(http.StatusOK, rd)

}
