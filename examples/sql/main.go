package main

import (
	"encoding/json"
	"fmt"

	"github.com/ra1n6ow/adming/internal/pkg/model"
	v1 "github.com/ra1n6ow/adming/pkg/api/system/v1"
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
	queryMenuList(db)
}

func expandChildren(db *gorm.DB) *gorm.DB {
	return db.Preload("Children", expandChildren)
}

func queryMenuList(db *gorm.DB) {
	// var user model.User
	// db.Preload("Role.Menus").Take(&user, "username = ?", "djf")
	// var menus []model.Menu
	// db.Preload("Menus").Find(&menus)
	// for _, role := range user.Roles {
	// 	fmt.Println(role.Menus) // 打印每个角色的详细信息
	// }
	// data, _ := json.Marshal(user.Role.Menus)
	// fmt.Println(string(data))

	// var menus []model.Menu
	// db.Preload("Children").Find(&user.Role.Menus, "parent_id is null")
	// data, _ := json.Marshal(user.Role.Menus)
	// fmt.Println(string(data))

	var user model.User
	db.Preload("Role").Find(&user, "username = ?", "admin")
	role := user.Role
	db.Preload("Menus", "parent_id is null").Preload("Menus.Children", expandChildren).Find(&role)
	// menus := role.Menus
	// db.Preload("Children", expandChildren).Find(&menus)

	menuList := v1.ConvertGetMenuListResponse(role.Menus)
	data, _ := json.Marshal(menuList)
	fmt.Println(string(data))
}
