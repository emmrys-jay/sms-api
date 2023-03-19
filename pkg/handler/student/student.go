package student

import (
	"github.com/Emmrys-Jay/altschool-sms/internal/model"
	studentService "github.com/Emmrys-Jay/altschool-sms/service/student"
	"github.com/Emmrys-Jay/altschool-sms/utility"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//	Create
//
// @Summary		add a student - Admin
// @Description	add a student - Admin
// @Tags			Student - Admin
// @Accept			json
// @Produce		json
// @Param			student	body		model.CreateStudent	true	"Student"
// @Success		201		{object}	utility.Response
// @Failure		400		{object}	utility.Response
// @Failure		401		{object}	utility.Response
// @Router			/student [post]
// @Security		JWTToken
func (base *Controller) Create(c *gin.Context) {
	var student model.CreateStudent

	if err := c.Bind(&student); err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "binding error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	if err := base.Validate.Struct(student); err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "validation error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	std, err := studentService.Create(c, &student)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "success", std)
	c.JSON(http.StatusCreated, rd)

}

//	Get
//
// @Summary		get a student
// @Description	get a student - Student
// @Tags			Student
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Student ID"
// @Success		200	{object}	utility.Response
// @Failure		400	{object}	utility.Response
// @Failure		401		{object}	utility.Response
// @Router			/student/{id} [get]
// @Security		JWTToken
func (base *Controller) Get(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	student, err := studentService.Get(c, uint(id))
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "success", student)
	c.JSON(http.StatusOK, rd)

}

//	List
//
// @Summary		list all students - Admin
// @Description	list all students - Admin
// @Tags			Student - Admin
// @Accept			json
// @Produce		json
// @Success		200	{object}	utility.Response
// @Failure		400	{object}	utility.Response
// @Failure		401		{object}	utility.Response
// @Router			/student [get]
// @Security		JWTToken
func (base *Controller) List(c *gin.Context) {

	students, err := studentService.List(c)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "success", students)
	c.JSON(http.StatusOK, rd)

}

//	ListStudentCourses
//
// @Summary		list courses taken by a student
// @Description	list courses taken by a student
// @Tags			Student
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Student ID"
// @Success		200	{object}	utility.Response
// @Failure		400	{object}	utility.Response
// @Failure		401		{object}	utility.Response
// @Router			/student/{id}/courses [get]
// @Security		JWTToken
func (base *Controller) ListStudentCourses(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	student, err := studentService.ListCourses(c, uint(id))
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "success", student)
	c.JSON(http.StatusOK, rd)
}

//	Update
//
// @Summary		update a student
// @Description	update a student
// @Tags			Student
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"Student ID"
// @Param			update	body		model.UpdateStudent	true	"Student Update"
// @Success		200		{object}	utility.Response
// @Failure		400		{object}	utility.Response
// @Failure		401		{object}	utility.Response
// @Router			/student/{id} [put]
// @Security		JWTToken
func (base *Controller) Update(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	var student model.UpdateStudent
	if err := c.Bind(&student); err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "binding error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	if err := base.Validate.Struct(student); err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "validation error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	if err := studentService.Update(c, uint(id), &student); err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, rd)

}

//	UpdateCourses
//
// @Summary		update or add single student courses - Admin
// @Description	update or add single student courses - Admin
// @Tags			Student - Admin
// @Accept			json
// @Produce		json
// @Param			id		path		int						true	"Student ID"
// @Param			update	body		[]model.CourseForDoc	true	"Student's courses'"
// @Success		200		{object}	utility.Response
// @Failure		400		{object}	utility.Response
// @Failure		401		{object}	utility.Response
// @Router			/student/{id}/courses [put]
// @Security		JWTToken
func (base *Controller) UpdateCourses(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	var courses []model.Course
	if err := c.Bind(&courses); err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "binding error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	if len(courses) == 0 {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "binding error", "no course specified", nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	if err := studentService.UpdateCourses(c, uint(id), courses); err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, rd)

}

//	Delete
//
// @Summary		delete a student - Admin
// @Description	delete a student - Admin
// @Tags			Student - Admin
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Student ID"
// @Success		200	{object}	utility.Response
// @Failure		400	{object}	utility.Response
// @Failure		401		{object}	utility.Response
// @Router			/student/{id} [delete]
// @Security		JWTToken
func (base *Controller) Delete(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	if err := studentService.Delete(c, uint(id)); err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, rd)

}
