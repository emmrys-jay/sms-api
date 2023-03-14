package course

import (
	"github.com/Emmrys-Jay/altschool-sms/internal/model"
	"github.com/Emmrys-Jay/altschool-sms/utility"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (base *Controller) Create(c *gin.Context) {
	var course model.Course

	if err := c.Bind(&course); err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "binding error", err, nil)
		c.JSON(http.StatusBadRequest, rd)
	}

	if err := base.Validate.Struct(course); err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "validation error", err, nil)
		c.JSON(http.StatusBadRequest, rd)
	}

	// TODO: Call service to create record

	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", gin.H{"user": "user object"})
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) Get(c *gin.Context) {

	idStr := c.Param("id")
	_, err := strconv.Atoi(idStr)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "validation error", err, nil)
		c.JSON(http.StatusBadRequest, rd)
	}

	// TODO: Call service to get record

	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", gin.H{"user": "login object"})
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) Update(c *gin.Context) {

	idStr := c.Param("id")
	_, err := strconv.Atoi(idStr)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "validation error", err, nil)
		c.JSON(http.StatusBadRequest, rd)
	}

	// TODO: Call service to update record

	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", gin.H{"user": "login object"})
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) Delete(c *gin.Context) {

	idStr := c.Param("id")
	_, err := strconv.Atoi(idStr)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "validation error", err, nil)
		c.JSON(http.StatusBadRequest, rd)
	}

	// TODO: Call service to delete record

	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", gin.H{"user": "login object"})
	c.JSON(http.StatusOK, rd)

}
