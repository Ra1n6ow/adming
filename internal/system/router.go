// Copyright (c) 2024 ra1n6ow <jeffduuu@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ra1n6ow/R-admin.

package system

import (
	"github.com/gin-gonic/gin"

	"github.com/ra1n6ow/adming/internal/pkg/core"
	"github.com/ra1n6ow/adming/internal/pkg/errno"
	"github.com/ra1n6ow/adming/internal/pkg/log"
	mw "github.com/ra1n6ow/adming/internal/pkg/middleware"
	"github.com/ra1n6ow/adming/internal/system/controller/v1/menu"
	"github.com/ra1n6ow/adming/internal/system/controller/v1/user"
	"github.com/ra1n6ow/adming/internal/system/store"
)

// installRouters 安装 miniblog 接口路由.
func installRouters(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// 注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz function called")

		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	g.GET("/getPermCode", func(c *gin.Context) {
		core.WriteResponse(c, nil, []string{"1000", "3000", "5000"})
	})

	uc := user.New(store.S)
	mc := menu.New(store.S)

	g.POST("/login", uc.Login)
	g.GET("/menuList", mw.Authn(), mc.GetMenuList)

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		// 创建 users 路由分组
		userv1 := v1.Group("/users")
		{
			userv1.POST("", uc.Create) // 创建用户
			userv1.Use(mw.Authn())
			userv1.GET(":name", uc.Get) // 获取用户详情
		}
	}

	return nil
}
