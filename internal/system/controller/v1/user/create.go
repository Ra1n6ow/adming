// Copyright (c) 2024 ra1n6ow <jeffduuu@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ra1n6ow/R-admin.

package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"github.com/ra1n6ow/adming/internal/pkg/core"
	"github.com/ra1n6ow/adming/internal/pkg/errno"
	"github.com/ra1n6ow/adming/internal/pkg/log"
	v1 "github.com/ra1n6ow/adming/pkg/api/system/v1"
)

// Create 创建一个新的用户.
func (ctrl *UserController) Create(c *gin.Context) {
	log.C(c).Infow("Create user function called")

	var r v1.CreateUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.ErrorResponse(c, errno.ErrBind)

		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.ErrorResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()))

		return
	}

	if err := ctrl.b.Users().Create(c, &r); err != nil {
		core.ErrorResponse(c, err)

		return
	}

	core.SuccessResponse(c, "Create user success", nil)
}
