package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/ra1n6ow/adming/internal/pkg/core"
	"github.com/ra1n6ow/adming/internal/pkg/log"
)

func (ctrl *MenuController) GetAll(c *gin.Context) {
	log.C(c).Infow("GetAll function called")

	menus, err := ctrl.b.Menus().GetAll(c)
	if err != nil {
		core.ErrorResponse(c, err)

		return
	}

	core.SuccessResponse(c, "Get all menus successful", menus)
}
