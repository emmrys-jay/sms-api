package course

import (
	"github.com/Emmrys-Jay/altschool-sms/internal/model"
	courseService "github.com/Emmrys-Jay/altschool-sms/service/course"
	"github.com/Emmrys-Jay/altschool-sms/utility"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Create godoc
//
//	@Summary		create a course
//	@Description	create a course
//	@Tags			Course
//	@Accept			json
//	@Produce		json
//	@Param			course	body		model.Course	true	"Course"
//	@Success		201		{object}	utility.Response
//	@Failure		400		{object}	utility.Response
//	@Router			/course [post]
func (base *Controller) Create(c *gin.Context) {
	var course model.Course

	if err := c.Bind(&course); err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "binding error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	if err := base.Validate.Struct(course); err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "validation error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	if err := courseService.Create(c, &course); err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "success", course)
	c.JSON(http.StatusCreated, rd)

}

// Get godoc
//
//	@Summary		get a course
//	@Description	get a course
//	@Tags			Course
//	@Accept			json
//	@Produce		json
//	@Param			code	path		string	true	"Course Code"
//	@Success		200		{object}	utility.Response
//	@Failure		400		{object}	utility.Response
//	@Router			/course/{id} [get]
func (base *Controller) Get(c *gin.Context) {

	code := strings.ToUpper(c.Param("code"))

	course, err := courseService.Get(c, code)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "success", course)
	c.JSON(http.StatusOK, rd)

}

// List godoc
//
//	@Summary		list all courses
//	@Description	list all courses
//	@Tags			Course
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utility.Response
//	@Failure		400	{object}	utility.Response
//	@Router			/course [get]
func (base *Controller) List(c *gin.Context) {

	courses, err := courseService.List(c)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "success", courses)
	c.JSON(http.StatusOK, rd)

}

// ListCourseStudents godoc
//
//	@Summary		list students taking a course
//	@Description	list students taking a course
//	@Tags			Course
//	@Accept			json
//	@Produce		json
//	@Param			code	path		string	true	"Course Code"
//	@Success		200		{object}	utility.Response
//	@Failure		400		{object}	utility.Response
//	@Router			/course/{id}/students [get]
func (base *Controller) ListCourseStudents(c *gin.Context) {

	code := strings.ToUpper(c.Param("code"))

	course, err := courseService.ListStudents(c, code)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "success", course)
	c.JSON(http.StatusOK, rd)

}

// Update godoc
//
//	@Summary		update a course
//	@Description	update a course
//	@Tags			Course
//	@Accept			json
//	@Produce		json
//	@Param			code	path		string			true	"Course Code"
//	@Param			update	body		model.Course	true	"Course Update"
//	@Success		200		{object}	utility.Response
//	@Failure		400		{object}	utility.Response
//	@Router			/course/{id} [put]
func (base *Controller) Update(c *gin.Context) {

	code := strings.ToUpper(c.Param("code"))

	var course model.Course
	if err := c.Bind(&course); err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "binding error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	if err := base.Validate.Struct(course); err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "validation error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	if err := courseService.Update(c, code, &course); err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, rd)

}

// Delete godoc
//
//	@Summary		delete a course
//	@Description	delete a course
//	@Tags			Course
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Course Code"
//	@Success		200	{object}	utility.Response
//	@Failure		400	{object}	utility.Response
//	@Router			/course/{id} [delete]
func (base *Controller) Delete(c *gin.Context) {

	code := strings.ToUpper(c.Param("code"))

	if err := courseService.Delete(c, code); err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, rd)

}
