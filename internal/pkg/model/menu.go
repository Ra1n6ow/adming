package model

import (
	"strconv"
	"strings"
	"time"
)

type Menu struct {
	ID              uint      `gorm:"column:id;primary_key;auto_increment"`
	ParentID        *uint     `gorm:"column:parent_id"`
	Type            string    `gorm:"column:type"`
	MenuName        string    `gorm:"column:menu_name"`
	Path            string    `gorm:"column:path"`
	Component       string    `gorm:"column:component"`
	Status          int       `gorm:"column:status"`
	Title           string    `gorm:"column:title"`
	Icon            string    `gorm:"column:icon"`
	OrderNo         int       `gorm:"column:order_no"`
	IsShow          string    `gorm:"column:is_show"`
	Permission      string    `gorm:"column:permission"`
	IgnoreKeepalive string    `gorm:"column:ignore_keepalive"`
	IsExt           string    `gorm:"column:is_ext"`
	Redirect        string    `gorm:"column:redirect"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
	Roles           []Role    `gorm:"many2many:role_menu"`
	Children        []Menu    `gorm:"foreignKey:ParentID;references:ID"`
}

// copier 的 CreateMenuRequest 钩子
func (menu *Menu) ParentMenu(parentid string) {
	var v string
	if strings.Contains(parentid, "-") {
		parts := strings.SplitN(parentid, "-", 3) // 最多分隔2次，即最多3个子串
		v = parts[len(parts)-1]
	} else {
		v = parentid
	}
	iv, _ := strconv.Atoi(v)
	ivp := uint(iv)
	menu.ParentID = &ivp
}
