package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/ra1n6ow/adming/internal/pkg/core"
	"github.com/ra1n6ow/adming/internal/pkg/log"
	"github.com/ra1n6ow/adming/pkg/token"
)

func (ctrl *MenuController) GetMeta(c *gin.Context) {
	log.C(c).Infow("GetMeta function called")

	// 从 token 中取出用户名
	identity, err := token.ParseRequest(c)
	if err != nil {
		core.ErrorResponse(c, err)
		return
	}
	menu, err := ctrl.b.Menus().GetMeta(c, identity)
	if err != nil {
		core.ErrorResponse(c, err)

		return
	}

	core.SuccessResponse(c, "Get menu meta for login success", menu)
}
