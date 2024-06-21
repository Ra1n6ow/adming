package menu

import (
	"github.com/ra1n6ow/adming/internal/system/biz"
	"github.com/ra1n6ow/adming/internal/system/store"
)

type MenuController struct {
	b biz.IBiz
}

// New 创建一个 user controller.
func New(ds store.IStore) *MenuController {
	return &MenuController{b: biz.NewBiz(ds)}
}
