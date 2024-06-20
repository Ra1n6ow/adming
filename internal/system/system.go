// Copyright (c) 2024 ra1n6ow <jeffduuu@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ra1n6ow/R-admin.

package system

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ra1n6ow/adming/internal/pkg/known"
	"github.com/ra1n6ow/adming/internal/pkg/log"
	mw "github.com/ra1n6ow/adming/internal/pkg/middleware"
	"github.com/ra1n6ow/adming/internal/pkg/model"
	"github.com/ra1n6ow/adming/internal/system/store"
	"github.com/ra1n6ow/adming/pkg/token"
	"github.com/ra1n6ow/adming/pkg/version/verflag"
)

var cfgFile string

// NewSystemCommand 创建一个 *cobra.Command 对象. 之后，可以使用 Command 对象的 Execute 方法来启动应用程序.
func NewSystemCommand() *cobra.Command {
	cmd := &cobra.Command{
		// 指定命令的名字，该名字会出现在帮助信息中
		Use: "system",
		// 命令的简短描述
		Short: "A Go systemtem admin",
		// 命令的详细描述
		Long: `A good Go admin project, used to manage systemtem.
Find more system information at:
        https://github.com/ra1n6ow/R-admin#readme`,
		// 命令出错时，不打印帮助信息。不需要打印帮助信息，设置为 true 可以保持命令出错时一眼就能看到错误信息
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行的 Run 函数，函数执行失败会返回错误信息
		RunE: func(cmd *cobra.Command, args []string) error {
			// 如果 `--version=true`，则打印版本并退出
			verflag.PrintAndExitIfRequested()

			// 初始化日志
			log.Init(logOptions())
			defer log.Sync() // Sync 将缓存中的日志刷新到磁盘文件中
			return run()
		},
		// 这里设置命令运行时，不需要指定命令行参数
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	// 以下设置，使得 initConfig 函数在每个命令运行时都会被调用以读取配置
	cobra.OnInitialize(initConfig)

	// 在这里您将定义标志和配置设置。

	// Cobra 支持持久性标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的每个子命令
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the system log configuration file. Empty string for no configuration file.")

	// Cobra 也支持本地标志，本地标志只能在其所绑定的命令上使用
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// 添加 --version 标志
	verflag.AddFlags(cmd.PersistentFlags())
	return cmd
}

// run 函数是实际的业务代码入口函数.
func run() error {
	// 打印所有的配置项及其值
	settings, _ := json.Marshal(viper.AllSettings())
	log.Infow(string(settings))

	// 打印 db -> username 配置项的值
	//log.Infow(viper.GetString("db.username"))
	// 初始化 store 层
	if err := initStore(); err != nil {
		return err
	}
	// migrateDB()
	// mockUserCreate()
	mockMenuCreate()

	// 设置 token 包的签发密钥，用于 token 包 token 的签发和解析
	token.Init(viper.GetString("jwt-secret"), known.XUsernameKey)

	// 设置 Gin 模式
	gin.SetMode(viper.GetString("runmode"))

	// 创建 Gin 引擎
	g := gin.New()

	// gin.Recovery() 中间件，用来捕获任何 panic，并恢复
	mws := []gin.HandlerFunc{gin.Recovery(), mw.NoCache, mw.Cors, mw.Secure, mw.RequestID()}

	g.Use(mws...)

	if err := installRouters(g); err != nil {
		return err
	}

	// 创建并运行 HTTP 服务器
	httpsrv := startInsecureServer(g)

	// 等待中断信号优雅地关闭服务器（10 秒超时)。
	//quit := make(chan os.Signal, 1)
	quit := make(chan os.Signal)
	// kill 默认会发送 systemcall.SIGTERM 信号
	// kill -2 发送 systemcall.SIGINT 信号，常用的 CTRL + C 就是触发系统 SIGINT 信号
	// kill -9 发送 systemcall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Infow("Shutting down server ...")

	// 创建 ctx 用于通知服务器 goroutine, 它有 10 秒时间完成当前正在处理的请求
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 10 秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过 10 秒就超时退出
	if err := httpsrv.Shutdown(ctx); err != nil {
		log.Errorw("Insecure Server forced to shutdown", "err", err)
		return err
	}

	log.Infow("Server exiting")
	return nil
}

// startInsecureServer 创建并运行 HTTP 服务器.
func startInsecureServer(g *gin.Engine) *http.Server {
	// 创建 HTTP Server 实例
	httpsrv := &http.Server{Addr: viper.GetString("addr"), Handler: g}

	// 运行 HTTP 服务器。在 goroutine 中启动服务器，它不会阻止下面的正常关闭处理流程
	// 打印一条日志，用来提示 HTTP 服务已经起来，方便排障
	log.Infow("Start to listening the incoming requests on http address", "addr", viper.GetString("addr"))
	go func() {
		if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalw(err.Error())
		}
	}()

	return httpsrv
}

func migrateDB() {
	db := store.S.DB()
	db.AutoMigrate(&model.Role{})
}
func mockUserCreate() {
	db := store.S.DB()
	var users []*model.User
	users = append(users, &model.User{
		Username: "admin",
		Password: "123456",
		Desc:     "系统管理员",
		HomePath: "/system/account",
		Roles: []*model.Role{
			{
				OrderNo:   1,
				RoleName:  "管理员",
				RoleValue: "admin",
				Status:    1,
			},
		},
	})

	users = append(users, &model.User{
		Username: "djf",
		Password: "123456",
		Desc:     "杜老二",
		HomePath: "/system/account",
		Roles: []*model.Role{
			{
				OrderNo:   2,
				RoleName:  "运维",
				RoleValue: "ops",
				Status:    1,
			},
		},
	})

	err := db.Create(&users).Error
	if err != nil {
		panic(err)
	}
	log.Infow("模拟创建用户完成")
}

func mockMenuCreate() {
	db := store.S.DB()
	var roles []*model.Role
	db.Find(&roles, "id = ?", 1)
	db.Create(&model.Menu{
		Name:        "系统管理",
		Title:       "系统管理",
		Path:        "/system",
		Component:   "/layouts/index",
		IsExt:       "0",
		IsShow:      "1",
		IsKeepalive: "1",
		Type:        "0",
		OrderNo:     1,
		Icon:        "setting",
		Roles:       roles,
	})
}
