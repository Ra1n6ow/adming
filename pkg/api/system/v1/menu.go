package v1

type GetMenuResponse struct {
	ID              string             `json:"id"`
	TreeID          string             `json:"treeID,omitempty"`
	Icon            string             `json:"icon,omitempty"`
	Component       string             `json:"component,omitempty"`
	Type            string             `json:"type,omitempty"`
	MenuName        string             `json:"menuName,omitempty"`
	Permission      string             `json:"permission,omitempty"`
	OrderNo         int                `json:"orderNo"`
	CreateTime      string             `json:"createTime,omitempty"`
	RoutePath       string             `json:"routePath,omitempty"`
	Redirect        string             `json:"redirect,omitempty"`
	Title           string             `json:"title,omitempty"`
	IgnoreKeepAlive bool               `json:"ignoreKeepAlive,omitempty"`
	Status          string             `json:"status,omitempty"`
	IsExt           string             `json:"isExt,omitempty"`
	Children        []*GetMenuResponse `json:"children,omitempty"`
}

type MenuMeta struct {
	Title           string `json:"title,omitempty"`
	Icon            string `json:"icon,omitempty"`
	IgnoreKeepAlive bool   `json:"ignoreKeepAlive,omitempty"`
	OrderNo         int    `json:"orderNo"`
}

type GetMenuMetaResponse struct {
	Path      string                 `json:"path,omitempty"`
	MenuName  string                 `json:"menuName,omitempty"`
	Component string                 `json:"component,omitempty"`
	Redirect  string                 `json:"redirect,omitempty"`
	Meta      MenuMeta               `json:"meta,omitempty"`
	Children  []*GetMenuMetaResponse `json:"children,omitempty"`
}

type CreateMenuRequest struct {
	Data CreateMenuRequestData `json:"data"`
}

type CreateMenuRequestData struct {
	MenuName        string `json:"menuName" valid:"alphanum,required,stringlength(1|255)"`
	OrderNo         int    `json:"orderNo" valid:"required"`
	Path            string `json:"routePath"`
	Icon            string `json:"icon"`
	Type            string `json:"type" valid:"required"`
	IsShow          string `json:"show"`
	Status          string `json:"status" valid:"required"`
	IsExt           string `json:"isExt"`
	Permission      string `json:"permission"`
	IgnoreKeepalive string `json:"keepalive"`
	ParentMenu      string `json:"parentMenu"`
	Component       string `json:"component"`
}
