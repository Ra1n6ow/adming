package model

import (
	"time"
)

type Role struct {
	ID        uint      `gorm:"column:id;primary_key;auto_increment"`
	OrderNo   int       `gorm:"column:order_no"`
	RoleName  string    `gorm:"column:role_name"`
	RoleValue string    `gorm:"column:role_value"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Desc      string    `gorm:"column:desc"`
	Status    int       `gorm:"column:status"`
	// Users     []User    `gorm:"many2many:user_role"`
	Menus []Menu `gorm:"many2many:role_menu"`
}

// // TableName sets the insert table name for this struct type
// func (r *Role) TableName() string {
// 	return "role"
// }
