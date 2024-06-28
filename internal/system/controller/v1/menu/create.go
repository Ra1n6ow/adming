package menu

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/ra1n6ow/adming/internal/pkg/core"
	"github.com/ra1n6ow/adming/internal/pkg/errno"
	"github.com/ra1n6ow/adming/internal/pkg/log"
	v1 "github.com/ra1n6ow/adming/pkg/api/system/v1"
)

func (ctrl *MenuController) Create(c *gin.Context) {
	log.C(c).Infow("Create menu function called")

	var r v1.CreateMenuRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.ErrorResponse(c, errno.ErrBind)

		return
	}

	if _, err := govalidator.ValidateStruct(r.Data); err != nil {
		core.ErrorResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()))

		return
	}

	if err := ctrl.b.Menus().Create(c, &r.Data); err != nil {
		core.ErrorResponse(c, err)

		return
	}

	core.SuccessResponse(c, "Create menu success", nil)
}
