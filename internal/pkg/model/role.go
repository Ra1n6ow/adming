package model

import (
	"time"
)

type Role struct {
	ID        uint      `gorm:"column:id;primary_key;auto_increment"` //
	OrderNo   int       `gorm:"column:orderNo"`                       //
	RoleName  string    `gorm:"column:roleName"`                      //
	RoleValue string    `gorm:"column:roleValue"`                     //
	CreatedAt time.Time `gorm:"column:createdAt"`                     //
	UpdatedAt time.Time `gorm:"column:updatedAt"`                     //
	Remark    string    `gorm:"column:remark"`                        //
	Status    int       `gorm:"column:status"`                        //
	Users     []*User   `gorm:"many2many:user_roles"`
}

// TableName sets the insert table name for this struct type
func (r *Role) TableName() string {
	return "roles"
}
