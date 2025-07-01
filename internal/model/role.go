package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `json:"name" gorm:"uniqueIndex"`
	User []User `gorm:"many2many:user_roles;" json:"-"`
}
