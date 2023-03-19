package health

import (
	"fmt"
	"net/http"

	"github.com/Emmrys-Jay/altschool-sms/internal/model"
	"github.com/Emmrys-Jay/altschool-sms/service/ping"
	"github.com/Emmrys-Jay/altschool-sms/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
	Validate *validator.Validate
	Logger   *utility.Logger
}

// Post godoc
//
//	@Summary		check api health
//	@Description	check api health
//	@Tags			Health
//	@Accept			json
//	@Produce		json
//	@Param			ping	body		model.Ping	true	"Ping"
//	@Success		200		{object}	utility.Response
//	@Failure		400		{object}	utility.Response
//	@Router			/health [post]
func (base *Controller) Post(c *gin.Context) {
	var (
		req = model.Ping{}
	)

	err := c.ShouldBind(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Failed to parse request body", err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	err = base.Validate.Struct(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Validation failed", utility.ValidationResponse(err, base.Validate), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}
	if !ping.ReturnTrue() {
		rd := utility.BuildErrorResponse(http.StatusInternalServerError, "error", "ping failed", fmt.Errorf("ping failed"), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}
	base.Logger.Info("ping successfull")

	rd := utility.BuildSuccessResponse(http.StatusOK, "ping successfull", req.Message)
	c.JSON(http.StatusOK, rd)

}

// Get godoc
//
//	@Summary		check api health
//	@Description	check api health
//	@Tags			Health
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utility.Response
//	@Failure		400	{object}	utility.Response
//	@Router			/health [get]
func (base *Controller) Get(c *gin.Context) {
	if !ping.ReturnTrue() {
		rd := utility.BuildErrorResponse(http.StatusInternalServerError, "error", "ping failed", fmt.Errorf("ping failed"), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}
	base.Logger.Info("ping successfull")
	rd := utility.BuildSuccessResponse(http.StatusOK, "ping successfull", gin.H{"user": "user object"})
	c.JSON(http.StatusOK, rd)

}
