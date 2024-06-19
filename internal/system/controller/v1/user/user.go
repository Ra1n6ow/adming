// Copyright (c) 2024 ra1n6ow <jeffduuu@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ra1n6ow/R-admin.

package user

import (
	"github.com/ra1n6ow/R-admin/internal/system/biz"
	"github.com/ra1n6ow/R-admin/internal/system/store"
)

// UserController 是 user 模块在 Controller 层的实现，用来处理用户模块的请求.
type UserController struct {
	b biz.IBiz
}

// New 创建一个 user controller.
func New(ds store.IStore) *UserController {
	return &UserController{b: biz.NewBiz(ds)}
}
