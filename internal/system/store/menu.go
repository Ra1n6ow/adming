package store

import (
	"context"

	"github.com/ra1n6ow/adming/internal/pkg/model"
	"gorm.io/gorm"
)

type MenuStore interface {
	getMenuByUser(ctx context.Context, username string) ([]*model.Menu, error)
}

type menus struct {
	db *gorm.DB
}

// 确保 users 实现了 UserStore 接口.
var _ MenuStore = (*menus)(nil)

func newMenus(db *gorm.DB) *menus {
	return &menus{db}
}

func (m *menus) getMenuByUser(ctx context.Context, username string) ([]*model.Menu, error) {
	var menus []*model.Menu
	return menus, nil
}
