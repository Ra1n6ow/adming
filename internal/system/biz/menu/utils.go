package menu

import (
	"fmt"
	"strconv"

	"github.com/ra1n6ow/adming/internal/pkg/model"
	v1 "github.com/ra1n6ow/adming/pkg/api/system/v1"
)

func uniqueMenus(menus []model.Menu) []model.Menu {
	uniqueMap := make(map[uint]bool)
	var uniqueList []model.Menu

	for _, menu := range menus {
		if _, exists := uniqueMap[menu.ID]; !exists {
			uniqueMap[menu.ID] = true
			uniqueList = append(uniqueList, menu)
		}
	}

	return uniqueList
}

func buildMenuTree(menus []model.Menu) []model.Menu {
	// 找到所有顶层节点
	var topLevelMenus []model.Menu
	for _, menu := range menus {
		if menu.ParentID == nil {
			topLevelMenus = append(topLevelMenus, menu)
		}
	}

	// 递归构建菜单树
	var result []model.Menu
	for _, topLevelMenu := range topLevelMenus {
		result = append(result, buildSubmenu(menus, &topLevelMenu))
	}

	return result
}

func buildSubmenu(menus []model.Menu, parent *model.Menu) model.Menu {
	var children []model.Menu
	for _, menu := range menus {
		if menu.ParentID != nil && *menu.ParentID == parent.ID {
			children = append(children, buildSubmenu(menus, &menu))
		}
	}

	parent.Children = children
	return *parent
}

func ConvertGetMenuMetaResponse(menus []model.Menu) []*v1.GetMenuMetaResponse {
	getMenuListResponses := make([]*v1.GetMenuMetaResponse, len(menus))
	for i, menu := range menus {
		getMenuListResponses[i] = &v1.GetMenuMetaResponse{
			Path:      menu.Path,
			MenuName:  menu.MenuName,
			Component: menu.Component,
			Redirect:  menu.Redirect,
			Meta: v1.MenuMeta{
				Title:           menu.Title,
				Icon:            menu.Icon,
				IgnoreKeepAlive: menu.IgnoreKeepalive == "1",
				OrderNo:         menu.OrderNo,
			},
			Children: ConvertGetMenuMetaResponse(menu.Children),
		}
	}
	return getMenuListResponses
}

var parentTreeID string

func ConvertGetMenuResponse(menus []model.Menu) []*v1.GetMenuResponse {
	if len(menus) > 0 && menus[0].ParentID != nil {
		parentTreeID += fmt.Sprintf("%d-", *menus[0].ParentID)
	}
	getMenuResponses := make([]*v1.GetMenuResponse, len(menus))
	for i, menu := range menus {
		if menu.ParentID == nil {
			parentTreeID = ""
		}
		getMenuResponses[i] = &v1.GetMenuResponse{
			ID:              strconv.Itoa(int(menu.ID)),
			TreeID:          fmt.Sprintf("%s%d", parentTreeID, menu.ID),
			Icon:            menu.Icon,
			Component:       menu.Component,
			Type:            menu.Type,
			MenuName:        menu.MenuName,
			Permission:      menu.Permission,
			OrderNo:         menu.OrderNo,
			RoutePath:       menu.Path,
			Status:          strconv.Itoa(int(menu.Status)),
			CreateTime:      menu.CreatedAt.Format("2006-01-02 15:04:05"),
			Redirect:        menu.Redirect,
			Title:           menu.Title,
			IgnoreKeepAlive: menu.IgnoreKeepalive == "1",
			Children:        ConvertGetMenuResponse(menu.Children),
		}
	}
	return getMenuResponses
}
