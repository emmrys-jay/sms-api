package course

import (
	"github.com/Emmrys-Jay/altschool-sms/internal/model"
	courseService "github.com/Emmrys-Jay/altschool-sms/service/course"
	"github.com/Emmrys-Jay/altschool-sms/utility"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Create
//
// @Summary		create a course - Admin
// @Description	create a course - Admin
// @Tags			Course - Admin
// @Accept			json
// @Produce		json
// @Param			course	body		model.CourseForDoc	true	"Course"
// @Success		201		{object}	utility.Response
// @Failure		400		{object}	utility.Response
// @Failure		401		{object}	utility.Response
// @Router			/course [post]
// @Security		JWTToken
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

//	Get
//
// @Summary		get a course - Admin
// @Description	get a course - Admin
// @Tags			Course - Admin
// @Accept			json
// @Produce		json
// @Param			code	path		string	true	"Course Code"
// @Success		200		{object}	utility.Response
// @Failure		400		{object}	utility.Response
// @Failure		401		{object}	utility.Response
// @Router			/course/{code} [get]
// @Security		JWTToken
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

// List
//
//	@Summary		list all courses - Admin
//	@Description	list all courses - Admin
//	@Tags			Course - Admin
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utility.Response
//	@Failure		400	{object}	utility.Response
//	@Failure		401		{object}	utility.Response
//	@Router			/course [get]
//	@Security		JWTToken
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

//	ListCourseStudents
//
// @Summary		list students taking a course - Admin
// @Description	list students taking a course - Admin
// @Tags			Course - Admin
// @Accept			json
// @Produce		json
// @Param			code	path		string	true	"Course Code"
// @Success		200		{object}	utility.Response
// @Failure		400		{object}	utility.Response
// @Failure		401		{object}	utility.Response
// @Router			/course/{code}/students [get]
// @Security		JWTToken
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

//	Update
//
// @Summary		update a course - Admin
// @Description	update a course - Admin
// @Tags			Course - Admin
// @Accept			json
// @Produce		json
// @Param			code	path		string				true	"Course Code"
// @Param			update	body		model.UpdateCourse	true	"Course Update"
// @Success		200		{object}	utility.Response
// @Failure		400		{object}	utility.Response
// @Failure		401		{object}	utility.Response
// @Router			/course/{code} [put]
// @Security		JWTToken
func (base *Controller) Update(c *gin.Context) {

	code := strings.ToUpper(c.Param("code"))

	var course model.UpdateCourse
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

//	Delete
//
// @Summary		delete a course - Admin
// @Description	delete a course - Admin
// @Tags			Course - Admin
// @Accept			json
// @Produce		json
// @Param			code	path		string	true	"Course Code"
// @Success		200	{object}	utility.Response
// @Failure		400	{object}	utility.Response
// @Failure		401		{object}	utility.Response
// @Router			/course/{code} [delete]
// @Security		JWTToken
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
