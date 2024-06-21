package v1

type GetMenuListResponse struct {
	Path      string `json:"path,omitempty"`
	Name      string `json:"name,omitempty"`
	Component string `json:"component,omitempty"`
	Redirect  string `json:"redirect,omitempty"`
	Meta      struct {
		Title           string `json:"title,omitempty"`
		Icon            string `json:"icon,omitempty"`
		IgnoreKeepAlive string `json:"ignoreKeepAlive,omitempty"`
	} `json:"meta"`
	Children []GetMenuListResponse `json:"children,omitempty"`
}
