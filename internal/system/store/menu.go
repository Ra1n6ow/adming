package store

import (
	"context"

	"github.com/ra1n6ow/adming/internal/pkg/model"
	"gorm.io/gorm"
)

type MenuStore interface {
	GetMenuList(ctx context.Context, username string) ([]model.Menu, error)
}

type menus struct {
	db *gorm.DB
}

// 确保 users 实现了 UserStore 接口.
var _ MenuStore = (*menus)(nil)

func newMenus(db *gorm.DB) *menus {
	return &menus{db}
}

func expandChildren(db *gorm.DB) *gorm.DB {
	return db.Preload("Children", expandChildren)
}
func (m *menus) GetMenuList(ctx context.Context, username string) ([]model.Menu, error) {
	var user model.User
	if err := m.db.Preload("Role").Find(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}
	role := user.Role
	if err := m.db.Preload("Menus", "parent_id is null").Preload("Menus.Children", expandChildren).Find(&role).Error; err != nil {
		return nil, err
	}

	return role.Menus, nil
}
