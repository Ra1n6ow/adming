package v1

// GetRoleResponse 指定了 `GET /v1/roles/{name}` 接口的返回参数.
type GetRoleResponse RoleInfo

// RoleInfo 指定了角色的详细信息.
type RoleInfo struct {
	OrderNo   int    `json:"orderNo"`
	RoleName  string `json:"roleName"`
	RoleValue string `json:"roleValue"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Remark    string `json:"remark"` //
	Status    int    `json:"status"`
}
