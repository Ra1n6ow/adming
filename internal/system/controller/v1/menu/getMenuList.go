package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/ra1n6ow/adming/internal/pkg/core"
	"github.com/ra1n6ow/adming/internal/pkg/log"
	"github.com/ra1n6ow/adming/pkg/token"
)

func (ctrl *MenuController) GetMenuList(c *gin.Context) {
	log.C(c).Infow("GetMenuList function called")

	// 从 token 中取出用户名
	identity, err := token.ParseRequest(c)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	menu, err := ctrl.b.Menus().GetMenuList(c, identity)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, menu)
}
