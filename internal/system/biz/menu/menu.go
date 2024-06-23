package menu

import (
	"context"
	"errors"

	"github.com/ra1n6ow/adming/internal/pkg/errno"
	"github.com/ra1n6ow/adming/internal/system/store"
	v1 "github.com/ra1n6ow/adming/pkg/api/system/v1"
	"gorm.io/gorm"
)

type MenuBiz interface {
	GetMenuList(ctx context.Context, username string) ([]*v1.GetMenuListResponse, error)
}

type menuBiz struct {
	ds store.IStore
}

var _ MenuBiz = (*menuBiz)(nil)

func New(ds store.IStore) *menuBiz {
	return &menuBiz{ds: ds}
}

func (b *menuBiz) GetMenuList(ctx context.Context, username string) ([]*v1.GetMenuListResponse, error) {
	menus, err := b.ds.Menus().GetMenuList(ctx, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrUserNotFound
		}
		return nil, err
	}
	resp := v1.ConvertGetMenuListResponse(menus)
	return resp, nil
}
