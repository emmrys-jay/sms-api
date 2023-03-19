package user

import (
	"net/http"

	"github.com/Emmrys-Jay/altschool-sms/internal/model"
	userService "github.com/Emmrys-Jay/altschool-sms/service/user"

	"github.com/Emmrys-Jay/altschool-sms/utility"
	"github.com/gin-gonic/gin"
)

//	Login
//
// @Summary		login as an admin or student
// @Description	login as an admin or student
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			login-info	body		model.Login	true	"Login Info"
// @Success		200			{object}	utility.Response
// @Failure		400			{object}	utility.Response
// @Router			/login [post]
func (base *Controller) Login(c *gin.Context) {

	var user model.Login
	if err := c.Bind(&user); err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	session, err := userService.Login(c, &user)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "failed", "error", err.Error(), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "success", session)
	c.JSON(http.StatusOK, rd)

}
