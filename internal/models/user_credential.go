package models

import (
	"time"

	"gorm.io/gorm"
)

type UserCredential struct {
	Email    string `gorm:"primaryKey;not null;index:,unique"`
	Password string

	User User `gorm:"foreignKey:Email;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
