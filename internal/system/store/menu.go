package store

import (
	"context"

	"github.com/ra1n6ow/adming/internal/pkg/model"
	"gorm.io/gorm"
)

type MenuStore interface {
	Create(ctx context.Context, menu *model.Menu) error
	GetMeta(ctx context.Context, username string) ([]model.Menu, error)
	GetAll(ctx context.Context) ([]model.Menu, error)
}

type menus struct {
	db *gorm.DB
}

// 确保 users 实现了 UserStore 接口.
var _ MenuStore = (*menus)(nil)

func newMenus(db *gorm.DB) *menus {
	return &menus{db}
}

func (m *menus) Create(ctx context.Context, menu *model.Menu) error {
	return m.db.Create(&menu).Error
}

func (m *menus) GetMeta(ctx context.Context, username string) ([]model.Menu, error) {
	var user model.User

	if err := m.db.Preload("Roles.Menus", func(db *gorm.DB) *gorm.DB {
		return db.Order("menu.order_no")
	}).Find(&user, "username = ?", "admin").Error; err != nil {
		return nil, err
	}
	var menuList []model.Menu
	for _, role := range user.Roles {
		menuList = append(menuList, role.Menus...)
	}

	return menuList, nil
}

func (m *menus) GetAll(ctx context.Context) ([]model.Menu, error) {
	var menus []model.Menu
	if err := m.db.Order("order_no").Find(&menus).Error; err != nil {
		return nil, err
	}

	return menus, nil
}
