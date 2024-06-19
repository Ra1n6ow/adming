// Copyright (c) 2024 ra1n6ow <jeffduuu@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ra1n6ow/R-admin.

package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ra1n6ow/R-admin/internal/pkg/core"
	"github.com/ra1n6ow/R-admin/internal/pkg/errno"
	"github.com/ra1n6ow/R-admin/internal/pkg/log"
	v1 "github.com/ra1n6ow/R-admin/pkg/api/system/v1"
)

// Login 登录 system 并返回一个 JWT Token.
func (ctrl *UserController) Login(c *gin.Context) {
	log.C(c).Infow("Login function called")

	var r v1.LoginRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	resp, err := ctrl.b.Users().Login(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, resp)
}
