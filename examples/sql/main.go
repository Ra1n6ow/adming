package main

import (
	"encoding/json"
	"fmt"

	"github.com/ra1n6ow/adming/internal/pkg/model"
	"github.com/ra1n6ow/adming/internal/system/biz/menu"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s`,
		"root",
		"123456",
		"127.0.0.1:3306",
		"adming",
		true,
		"UTC")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}
	// queryMenuMeta(db)
	// quaryMenus(db)
	menuCreate(db)
}

func expandChildren(db *gorm.DB) *gorm.DB {
	return db.Preload("Children", expandChildren)
}

func queryMenuMeta(db *gorm.DB) {
	// var user model.User
	// db.Preload("Role.Menus").Take(&user, "username = ?", "djf")
	// var menus []model.Menu
	// db.Preload("Menus").Find(&menus)
	// for _, role := range user.Roles {
	// 	fmt.Println(role.Menus) // 打印每个角色的详细信息
	// }
	// data, _ := json.Marshal(user.Role.Menus)
	// fmt.Println(string(data))

	// menuList := v1.ConvertGetMenuListResponse(role.Menus)
	// data, _ := json.Marshal(menuList)
	// fmt.Println(string(data))

	var user model.User
	db.Preload("Roles.Menus", func(db *gorm.DB) *gorm.DB {
		return db.Order("menu.order_no")
	}).Find(&user, "username = ?", "admin")
	// roles := user.Roles
	// db.Preload("Menus").Find(&roles)
	var menuList []model.Menu
	for _, role := range user.Roles {
		menuList = append(menuList, role.Menus...)
	}
	uniqueList := uniqueMenus(menuList)
	rootMenus := buildMenuTree(uniqueList)
	// result := menu.ConvertGetMenuMetaResponse(rootMenus)
	// data, _ := json.Marshal(result)
	// fmt.Println(string(data))
	// fmt.Println(user.Roles)
	data, _ := json.Marshal(rootMenus)
	fmt.Println(string(data))
}

func uniqueMenus(menus []model.Menu) []model.Menu {
	uniqueMap := make(map[uint]bool)
	var uniqueList []model.Menu

	for _, menu := range menus {
		if _, exists := uniqueMap[menu.ID]; !exists {
			uniqueMap[menu.ID] = true
			uniqueList = append(uniqueList, menu)
		}
	}

	return uniqueList
}

func buildMenuTree(menus []model.Menu) []model.Menu {
	// 找到所有顶层节点
	var topLevelMenus []model.Menu
	for _, menu := range menus {
		if menu.ParentID == nil {
			topLevelMenus = append(topLevelMenus, menu)
		}
	}

	// 递归构建菜单树
	var result []model.Menu
	for _, topLevelMenu := range topLevelMenus {
		result = append(result, buildSubmenu(menus, &topLevelMenu))
	}

	return result
}

func buildSubmenu(menus []model.Menu, parent *model.Menu) model.Menu {
	var children []model.Menu
	for _, menu := range menus {
		if menu.ParentID != nil && *menu.ParentID == parent.ID {
			children = append(children, buildSubmenu(menus, &menu))
		}
	}

	parent.Children = children
	return *parent
}

func quaryMenus(db *gorm.DB) {
	// var user model.User
	// db.Preload("Roles.Menus").Find(&user, "username = ?", "admin")
	// role := user.Roles[0]
	// db.Preload("Menus").Find(&role)
	// var menuList []model.Menu
	// for _, role := range user.Roles {
	// 	menuList = append(menuList, role.Menus...)
	// }
	// uniqueList := uniqueMenus(menuList)
	var menus []model.Menu
	db.Order("order_no").Find(&menus)
	rootMenus := buildMenuTree(menus)
	res := menu.ConvertGetMenuResponse(rootMenus)
	data, _ := json.Marshal(res)
	fmt.Println(string(data))
}

func menuCreate(db *gorm.DB) {
}
