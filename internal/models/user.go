package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Email string `gorm:"unique;index:,unique"`
	Name  string
	Role  string

	Client []Client `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
