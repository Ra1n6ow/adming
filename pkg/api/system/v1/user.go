// Copyright (c) 2024 ra1n6ow <jeffduuu@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ra1n6ow/R-admin.

package v1

// LoginRequest 指定了 `POST /login` 接口的请求参数.
type LoginRequest struct {
	Username string `json:"username" valid:"alphanum,required,stringlength(1|255)"`
	Password string `json:"password" valid:"required,stringlength(6|18)"`
}

// LoginResponse 指定了 `POST /login` 接口的返回参数.
type LoginResponse struct {
	Token string `json:"token"`
}

// CreateUserRequest 指定了 `POST /v1/users` 接口的请求参数.
type CreateUserRequest struct {
	Username string `json:"username" valid:"alphanum,required,stringlength(1|255)"`
	Password string `json:"password" valid:"required,stringlength(6|18)"`
	Nickname string `json:"nickname" valid:"required,stringlength(1|255)"`
	Email    string `json:"email" valid:"required,email"`
	Phone    string `json:"phone" valid:"required,stringlength(11|11)"`
}

// GetUserResponse 指定了 `GET /v1/users/{name}` 接口的返回参数.
type GetUserResponse UserInfo

// UserInfo 指定了用户的详细信息.
type UserInfo struct {
	Username  string      `json:"username"`
	Avatar    string      `json:"avatar"`
	Desc      string      `json:"desc"`
	CreatedAt string      `json:"createdAt"`
	UpdatedAt string      `json:"updatedAt"`
	HomePath  string      `json:"homePath"` //
	Status    int         `json:"status"`
	Roles     []*RoleInfo `gorm:"many2many:user_roles"`
}
