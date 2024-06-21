package menu

import (
	"context"

	"github.com/ra1n6ow/adming/internal/system/store"
)

type MenuBiz interface {
}

type menuBiz struct {
	ds store.IStore
}

var _ MenuBiz = (*menuBiz)(nil)

func New(ds store.IStore) *menuBiz {
	return &menuBiz{ds: ds}
}

func (b *menuBiz) Get(ctx context.Context, username string) {

}
