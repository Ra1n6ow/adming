package model

import (
	"gorm.io/gorm"
	"time"

	"github.com/ra1n6ow/R-admin/pkg/auth"
)

type User struct {
	ID        uint      `gorm:"column:id;primary_key;auto_increment"` //
	Username  string    `gorm:"column:username"`                      //
	Password  string    `gorm:"column:password"`                      //
	Avatar    string    `gorm:"column:avatar"`                        //
	Desc      string    `gorm:"column:desc"`                          //
	HomePath  string    `gorm:"column:homePath"`                      //
	CreatedAt time.Time `gorm:"column:createdAt"`                     //
	UpdatedAt time.Time `gorm:"column:updatedAt"`                     //
	Status    int       `gorm:"column:status"`
	Roles     []*Role   `gorm:"many2many:user_roles"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "users"
}

// BeforeCreate 在创建数据库记录之前加密明文密码.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Encrypt the user password.
	u.Password, err = auth.Encrypt(u.Password)
	if err != nil {
		return err
	}

	return nil
}
