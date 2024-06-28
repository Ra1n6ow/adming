package menu

import (
	"context"
	"errors"
	"regexp"

	"github.com/jinzhu/copier"
	"github.com/ra1n6ow/adming/internal/pkg/errno"
	"github.com/ra1n6ow/adming/internal/pkg/model"
	"github.com/ra1n6ow/adming/internal/system/store"
	v1 "github.com/ra1n6ow/adming/pkg/api/system/v1"
	"gorm.io/gorm"
)

type MenuBiz interface {
	GetMeta(ctx context.Context, username string) ([]*v1.GetMenuMetaResponse, error)
	GetAll(ctx context.Context) ([]*v1.GetMenuResponse, error)
	Create(ctx context.Context, r *v1.CreateMenuRequestData) error
}

type menuBiz struct {
	ds store.IStore
}

var _ MenuBiz = (*menuBiz)(nil)

func New(ds store.IStore) *menuBiz {
	return &menuBiz{ds: ds}
}

func (b *menuBiz) Create(ctx context.Context, r *v1.CreateMenuRequestData) error {
	var menu model.Menu
	_ = copier.Copy(&menu, &r)
	// 外键约束：当没有父级时，传空
	if *menu.ParentID == 0 {
		menu.ParentID = nil
	}

	if err := b.ds.Menus().Create(ctx, &menu); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key '.*'", err.Error()); match {
			return errno.ErrMenuAlreadyExist
		}
		if match, _ := regexp.MatchString(".*a foreign key constraint fails.*", err.Error()); match {
			return errno.ErrMenuConstraintFail
		}

		return err
	}

	return nil
}

func (b *menuBiz) GetMeta(ctx context.Context, username string) ([]*v1.GetMenuMetaResponse, error) {
	menus, err := b.ds.Menus().GetMeta(ctx, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrUserNotFound
		}
		return nil, err
	}
	uniqueList := uniqueMenus(menus)
	rootMenus := buildMenuTree(uniqueList)
	resp := ConvertGetMenuMetaResponse(rootMenus)
	return resp, nil
}

func (b *menuBiz) GetAll(ctx context.Context) ([]*v1.GetMenuResponse, error) {

	menus, err := b.ds.Menus().GetAll(ctx)
	if err != nil {
		return nil, err
	}
	rootMenus := buildMenuTree(menus)
	resp := ConvertGetMenuResponse(rootMenus)
	return resp, nil
}
