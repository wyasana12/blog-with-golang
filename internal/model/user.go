package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string     `json:"name"`
	Email     string     `json:"email" gorm:"uniqueIndex"`
	Password  string     `json:"-"`
	OTPToken  string     `json:"otp_token" gorm:"index"`
	OTPExpiry *time.Time `json:"otp_expiry"`
	Roles     []Role     `gorm:"many2many:user_roles;" json:"roles"`
}
