package v1

import (
	"github.com/ra1n6ow/adming/internal/pkg/model"
)

type MenuMeta struct {
	Title           string `json:"title,omitempty"`
	Icon            string `json:"icon,omitempty"`
	IgnoreKeepAlive bool   `json:"ignoreKeepAlive,omitempty"`
}

type GetMenuListResponse struct {
	Path      string                 `json:"path,omitempty"`
	Name      string                 `json:"name,omitempty"`
	Component string                 `json:"component,omitempty"`
	Redirect  string                 `json:"redirect,omitempty"`
	Meta      MenuMeta               `json:"meta,omitempty"`
	Children  []*GetMenuListResponse `json:"children,omitempty"`
}

// 将model.Menu转换为GetMenuListResponse
func ConvertGetMenuListResponse(menus []model.Menu) []*GetMenuListResponse {
	getMenuListResponses := make([]*GetMenuListResponse, len(menus))
	for i, menu := range menus {
		getMenuListResponses[i] = &GetMenuListResponse{
			Path:      menu.Path,
			Name:      menu.Name,
			Component: menu.Component,
			Redirect:  menu.Redirect,
			Meta: MenuMeta{
				Title:           menu.Title,
				Icon:            menu.Icon,
				IgnoreKeepAlive: menu.IgnoreKeepalive == "1",
			},
			Children: ConvertGetMenuListResponse(menu.Children),
		}
	}
	return getMenuListResponses
}
