package model

import "time"

type Menu struct {
	ID              uint      `gorm:"column:id;primary_key;auto_increment"`
	Pid             string    `gorm:"column:pid"`
	Type            string    `gorm:"column:type"`
	Name            string    `gorm:"column:name"`
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
	Roles           []*Role   `gorm:"many2many:role_menu"`
}
