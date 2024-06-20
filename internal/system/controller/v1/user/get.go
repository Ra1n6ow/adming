// Copyright (c) 2024 ra1n6ow <jeffduuu@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ra1n6ow/R-admin.

package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ra1n6ow/adming/internal/pkg/core"
	"github.com/ra1n6ow/adming/internal/pkg/log"
	v1 "github.com/ra1n6ow/adming/pkg/api/system/v1"
	"github.com/ra1n6ow/adming/pkg/token"
)

// Get 获取一个用户的详细信息.
func (ctrl *UserController) Get(c *gin.Context) {
	log.C(c).Infow("Get user function called")

	var user *v1.GetUserResponse
	var err error
	var identity string

	// 当请求为 /v1/users/:name 时，从 Param 中取出查询用户
	name := c.Param("name")
	// 当请求为 /v1/users/userinfo 时，从 token 中取出查询用户
	if name == "userinfo" {
		identity, err = token.ParseRequest(c)
		if err != nil {
			core.WriteResponse(c, err, nil)
			return
		}
	}

	user, err = ctrl.b.Users().Get(c, identity)

	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, user)
}
