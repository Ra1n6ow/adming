// Copyright (c) 2024 ra1n6ow <jeffduuu@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ra1n6ow/R-admin.

package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ra1n6ow/adming/internal/pkg/core"
	"github.com/ra1n6ow/adming/internal/pkg/errno"
	"github.com/ra1n6ow/adming/internal/pkg/log"
	v1 "github.com/ra1n6ow/adming/pkg/api/system/v1"
)

// Login 登录 system 并返回一个 JWT Token.
func (ctrl *UserController) Login(c *gin.Context) {
	log.C(c).Infow("Login function called")

	var r v1.LoginRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.ErrorResponse(c, errno.ErrBind)

		return
	}

	resp, err := ctrl.b.Users().Login(c, &r)
	if err != nil {
		core.ErrorResponse(c, err)

		return
	}

	core.SuccessResponse(c, "Login success", resp)
}
